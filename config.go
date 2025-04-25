package sqlbuilder

/*func NewCache(opts ...SelectOption) (*SelectConfig, error) {
	conf := NewSelectConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return nil, err
		}
	}
	return conf, nil
}*/

import (
	"fmt"
	"strings"
)

type SQLBody struct {
	text   map[string]string
	dbType string
}

func (s *SQLBody) getBindMark(i int) string {
	if strings.EqualFold(s.dbType, "mysql") {
		return fmt.Sprintf("%s", "?")
	} else if strings.EqualFold(s.dbType, "oracle") {
		return fmt.Sprintf(":%v", i)
	}
	return "?"
}

func (s *SQLBody) getWithPage(v string, i int) (t string) {
	var whereSQL strings.Builder

	if strings.EqualFold(s.dbType, "mysql") {
		whereSQL.WriteString(fmt.Sprintf(`select * from (`))
		whereSQL.WriteString(v)
		whereSQL.WriteString(") c limit ?,?")
	} else if strings.EqualFold(s.dbType, "oracle") {
		whereSQL.WriteString(v)
		whereSQL.WriteString(fmt.Sprintf(" offset :%v rows fetch next :%v rows only", i, i+1))
	}
	return whereSQL.String()
}

func (s *SQLBody) getWithSet(v string, dt string, set []string) (t string) {
	var whereSQL strings.Builder
	for i, col := range set {
		if i == 0 {
			whereSQL.WriteString(fmt.Sprintf("%v=%v", col, s.getBindMark(i+1)))
		} else {
			whereSQL.WriteString(fmt.Sprintf(",%v=%v", col, s.getBindMark(i+1)))
		}
	}
	t = strings.ReplaceAll(v, setColumn, whereSQL.String())
	return t
}

func (s *SQLBody) getWithWhereUnknown(v string, u string) (t string) {
	if len(u) > 0 {
		return strings.ReplaceAll(v, unknownWhere, u)
	}
	return strings.ReplaceAll(v, unknownWhere, "")
}

func (s *SQLBody) getWithWhereNull(v string, n []*NullColumn) (t string) {
	if len(n) > 0 {
		var whereSQL strings.Builder
		for _, column := range n {
			if column.IsNull {
				whereSQL.WriteString(fmt.Sprintf("%s %s is null ", column.Operator, column.ColumnName))
			} else {
				whereSQL.WriteString(fmt.Sprintf("%s %s is not null ", column.Operator, column.ColumnName))
			}

		}
		return strings.ReplaceAll(v, whereNullCol, whereSQL.String())
	}
	return strings.ReplaceAll(v, whereNullCol, "")
}

/*
*
第1参数: sql文本
第2参数: 数据库类型
第3参数: where条件
第4参数: n是绑定变量起始位. 例如update set 占了1~5,那么where就必须从6开始
*/
func (s *SQLBody) getWithWhere(v string, p []*PredicateColumn, n int) (t string) {
	var whereSQL strings.Builder
	if len(p) == 0 {
		t = strings.ReplaceAll(v, where, "1=1")
	} else {
		whereSQL.WriteString("1=1")
		var inCount int
		var inVal strings.Builder
		for i, col := range p {
			switch col.Expression {
			case "in":
				for ii := 0; ii < col.mustValueCount; ii++ {
					if ii == 0 {
						inVal.WriteString(fmt.Sprintf("%v", s.getBindMark(n+ii+i)))
					} else {
						inVal.WriteString(fmt.Sprintf(",%v", s.getBindMark(n+ii+i)))
					}

				}
				inCount = col.mustValueCount - 1
				whereSQL.WriteString(fmt.Sprintf(" %s %s %s (%v)", col.Operator, col.ColumnName, col.Expression, inVal.String()))

			default:
				whereSQL.WriteString(fmt.Sprintf(" %s %s %s %v", col.Operator, col.ColumnName, col.Expression, s.getBindMark(n+i+inCount))) // 开始位置+当前位置+增长位
			}
		}
		t = strings.ReplaceAll(v, where, whereSQL.String())
	}
	return t
}

func (s *SQLBody) getWithResult(v string, r []*ResultColumn) (t string) {
	var selectSQL strings.Builder
	for i, res := range r {
		if i == 0 {
			if len(res.ColumnName) > 0 {
				if len(res.Alias) > 0 {
					selectSQL.WriteString(fmt.Sprintf("%v %v", res.ColumnName, res.Alias))
				} else {
					selectSQL.WriteString(fmt.Sprintf("%v", res.ColumnName))
				}

			}
		} else {
			if len(res.ColumnName) > 0 {
				if len(res.Alias) > 0 {
					selectSQL.WriteString(fmt.Sprintf(",%v %v", res.ColumnName, res.Alias))
				} else {
					selectSQL.WriteString(fmt.Sprintf(",%v", res.ColumnName))
				}
			}
		}
	}
	t = strings.ReplaceAll(v, result, selectSQL.String())
	return t
}

func (s *SQLBody) getWithOrder(v string, o []*OrderColumn) (t string) {
	var orderSQL strings.Builder
	if len(o) > 0 {
		orderSQL.WriteString("order by ")
		for i, col := range o {
			if i == 0 {
				if len(col.ColumnName) > 0 {
					if len(col.Order) > 0 {
						orderSQL.WriteString(fmt.Sprintf("%s %s", col.ColumnName, col.Order))
					} else {
						orderSQL.WriteString(fmt.Sprintf("%s", col.ColumnName))
					}

				}
			} else {
				if len(col.ColumnName) > 0 {
					if len(col.Order) > 0 {
						orderSQL.WriteString(fmt.Sprintf(",%s %s", col.ColumnName, col.Order))
					} else {
						orderSQL.WriteString(fmt.Sprintf(",%s", col.ColumnName))
					}
				}
			}

		}
		t = strings.ReplaceAll(v, sort, orderSQL.String())
	} else {
		t = strings.ReplaceAll(v, sort, "")
	}
	return t
}

func (s *SQLBody) getWithColumn(v string, c []string) (t string) {
	var column strings.Builder
	var values strings.Builder
	for i, col := range c {
		if i == 0 {
			column.WriteString(fmt.Sprintf("%v", col))
			values.WriteString(fmt.Sprintf("%v", s.getBindMark(i+1)))
		} else {
			column.WriteString(fmt.Sprintf(",%v", col))
			values.WriteString(fmt.Sprintf(",%v", s.getBindMark(i+1)))
		}

	}
	v = strings.ReplaceAll(v, setColumn, column.String())
	v = strings.ReplaceAll(v, valueColumn, values.String())
	return v
}

func (s *SQLBody) getSelectForDatabase(k string, dt string, r []*ResultColumn, p []*PredicateColumn, n []*NullColumn, o []*OrderColumn, page bool, u string) (t string, e error) {
	v, ok := s.text[k]
	if !ok {
		return "", fmt.Errorf("sql text %v key does not exist", k)
	}

	if len(r) == 0 {
		return "", fmt.Errorf("sql key %s select result column must be specific", k)
	}

	// 生成结果列
	v = s.getWithResult(v, r)

	// 生成微词条件语句
	v = s.getWithWhere(v, p, 1)

	// 生成谓词条件列是null的语句
	v = s.getWithWhereNull(v, n)

	// 生成未知语句
	v = s.getWithWhereUnknown(v, u)

	// 生成排序语句
	v = s.getWithOrder(v, o)

	if page {
		v = s.getWithPage(v, len(p)+1)
	}

	return strings.TrimSpace(v), nil
}

func (s *SQLBody) getInsertForDatabase(k string, c []string) (t string, e error) {
	v, ok := s.text[k]
	if !ok {
		return "", fmt.Errorf("sql text %v key does not exist", k)
	}

	if len(c) == 0 {
		return "", fmt.Errorf("sql key %v Column must be set", k)
	}

	// 生成insert
	v = s.getWithColumn(v, c)

	return strings.TrimSpace(v), nil
}

func (s *SQLBody) getCountForDatabase(k string, p ...*PredicateColumn) (t string, e error) {
	v, ok := s.text[k]
	if !ok {
		return "", fmt.Errorf("sql text %v key does not exist", k)
	}
	// 生成where条件语句
	t = s.getWithWhere(v, p, 1)
	// 去除排序
	t = strings.ReplaceAll(t, sort, "")

	// 生成未知语句. count 暂时不支持负责的where语句
	// 去除未知where
	t = strings.ReplaceAll(t, unknownWhere, "")

	// 去除结果列
	t = strings.ReplaceAll(t, result, "count(1)")
	return t, nil
}

func (s *SQLBody) getDeleteForDatabase(k string, dt string, p []*PredicateColumn) (t string, e error) {
	v, ok := s.text[k]
	if !ok {
		return "", fmt.Errorf("sql text %v key does not exist", k)
	}

	// 生成谓词条件语句
	if len(p) > 0 {
		v = s.getWithWhere(v, p, 1)
	}

	return strings.TrimSpace(v), nil
}

func (s *SQLBody) getUpdateForDatabase(k string, dt string, set []string, p []*PredicateColumn) (t string, e error) {
	v, ok := s.text[k]
	if !ok {
		return "", fmt.Errorf("sql text %v key does not exist", k)
	}

	if len(set) == 0 {
		return "", fmt.Errorf("sql key %s select result column must be specific", k)
	}

	// 生成set值语句. 长度取set列个数的
	if len(set) > 0 {
		v = s.getWithSet(v, dt, set)
	}

	// 生成谓词条件语句. where绑定编码需要从set列个数开始
	if len(p) > 0 {
		v = s.getWithWhere(v, p, len(set)+1)
	}

	return strings.TrimSpace(v), nil
}

func (s *SQLBody) GetCount(k string, opts ...CountOption) (string, error) {
	conf := NewCountConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", err
		}
	}

	return s.getCountForDatabase(k, conf.predicate...)
}

func (s *SQLBody) GetPublicCount(tableName string, opts ...CountOption) (string, error) {
	conf := NewCountConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", err
		}
	}

	str, err := s.getCountForDatabase(InternalSelect, conf.predicate...)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(str, internalTable, tableName), nil
}

func (s *SQLBody) GetSelect(k string, opts ...SelectOption) (string, error) {
	conf := NewSelectConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", k, err)
		}
	}

	return s.getSelectForDatabase(k, s.dbType, conf.result, conf.predicate, conf.isNullCol, conf.order, conf.Page, conf.unknown)
}

func (s *SQLBody) GetPublicSelect(tableName string, opts ...SelectOption) (string, error) {
	conf := NewSelectConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", InternalSelect, err)
		}
	}

	str, err := s.getSelectForDatabase(InternalSelect, s.dbType, conf.result, conf.predicate, conf.isNullCol, conf.order, conf.Page, conf.unknown)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(str, internalTable, tableName), nil
}

func (s *SQLBody) GetUpdate(k string, opts ...UpdateOption) (string, error) {
	conf := NewUpdateConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", k, err)
		}
	}

	return s.getUpdateForDatabase(k, s.dbType, conf.setColumnWithUpdate, conf.predicate)
}

func (s *SQLBody) GetPublicUpdate(tableName string, opts ...UpdateOption) (string, error) {
	conf := NewUpdateConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", InternalUpdate, err)
		}
	}

	str, err := s.getUpdateForDatabase(InternalUpdate, s.dbType, conf.setColumnWithUpdate, conf.predicate)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(str, internalTable, tableName), nil
}

func (s *SQLBody) GetDelete(k string, opts ...DeleteOption) (string, error) {
	conf := NewDeleteConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", k, err)
		}
	}

	return s.getDeleteForDatabase(k, s.dbType, conf.predicate)
}

func (s *SQLBody) GetPublicDelete(tableName string, opts ...DeleteOption) (string, error) {
	conf := NewDeleteConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", InternalDelete, err)
		}
	}
	str, err := s.getDeleteForDatabase(InternalDelete, s.dbType, conf.predicate)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(str, internalTable, tableName), nil
}

func (s *SQLBody) GetInsert(k string, opts ...InsertOption) (string, error) {
	conf := NewInsertConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", k, err)
		}
	}

	return s.getInsertForDatabase(k, conf.columnWithInsert)
}

func (s *SQLBody) GetPublicInsert(tableName string, opts ...InsertOption) (string, error) {
	conf := NewInsertConfig()
	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return "", fmt.Errorf("sql key %v. error: %v", InternalInsert, err)
		}
	}

	str, err := s.getInsertForDatabase(InternalInsert, conf.columnWithInsert)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(str, internalTable, tableName), nil
}

func (s *SQLBody) InitSQLText(key, text string) {
	s.text[key] = text
}

func (s *SQLBody) SetTypeForMySQL() {
	s.dbType = "mysql"
}

func (s *SQLBody) SetTypeForOracle() {
	s.dbType = "oracle"
}

func (s *SQLBody) GetDBType() string {
	return s.dbType
}

func NewSQLBody() *SQLBody {
	s := new(SQLBody)
	s.text = make(map[string]string)
	s.SetTypeForMySQL()
	s.InitSQLText(InternalInsert, InternalInsertText)
	s.InitSQLText(InternalDelete, InternalDeleteText)
	s.InitSQLText(InternalUpdate, InternalUpdateText)
	s.InitSQLText(InternalSelect, InternalSelectText)
	return s
}
