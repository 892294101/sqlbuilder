package sqlbuilder

import (
	"fmt"
	"strings"
)

type ResultColumn struct {
	ColumnName string // 列名
	Alias      string
}

type PredicateColumn struct {
	Operator       string // 逻辑操作符and、or
	ColumnName     string // 列名
	Expression     string // 表达式
	mustValueCount int    // 多指数量
}

type OrderColumn struct {
	ColumnName string // 列名
	Order      string // 排序. 不为空则排序. 但是排序必须为asc desc
}

type InsertConfig struct {
	columnWithInsert []string // insert 列名数据
}

type DeleteConfig struct {
	predicate []*PredicateColumn // delete where条件数据
}

type UpdateConfig struct {
	setColumnWithUpdate []string           // update set列数据
	predicate           []*PredicateColumn // update where条件数据
}

type SelectConfig struct {
	result    []*ResultColumn    // select 返回列
	predicate []*PredicateColumn // select where条件数据
	order     []*OrderColumn     // select 排序, 必须不可以是count
	unknown   string             // 用于未知的where条件. 如复杂的where条件, 而本函数实现不了. 则可以使用未知where条件来实现.
	Page      bool               // select 是否分页
}

type CountConfig struct {
	predicate []*PredicateColumn // select where条件数据
}

type InsertOption func(*InsertConfig) error
type DeleteOption func(*DeleteConfig) error
type UpdateOption func(*UpdateConfig) error
type CountOption func(*CountConfig) error
type SelectOption func(*SelectConfig) error

func NewInsertConfig() *InsertConfig { return new(InsertConfig) }
func NewDeleteConfig() *DeleteConfig { return new(DeleteConfig) }
func NewUpdateConfig() *UpdateConfig { return new(UpdateConfig) }
func NewCountConfig() *CountConfig   { return new(CountConfig) }
func NewSelectConfig() *SelectConfig { return new(SelectConfig) }

// SetSelectWhereColumn 设置查询where
func SetSelectWhereColumn(Operator, ColumnName, Expression string) SelectOption {
	return func(c *SelectConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 || len(Expression) == 0 || strings.EqualFold(strings.ToLower(Expression), "in") {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: Expression})
		return nil
	}
}

// SetSelectWhereUnknown 设置查询where
func SetSelectWhereUnknown(Expression string) SelectOption {
	return func(c *SelectConfig) error {
		if len(Expression) > 0 {
			c.unknown = Expression
		}
		return nil
	}
}

// SetSelectResultColumn 设置查询返回列
func SetSelectResultColumn(col string, a ...string) SelectOption {
	return func(c *SelectConfig) error {
		if len(col) > 0 {
			switch {
			case len(a) > 1:
				return fmt.Errorf("only one column alias can be set for column %v", col)
			case len(a) == 1:
				c.result = append(c.result, &ResultColumn{ColumnName: strings.TrimSpace(col), Alias: strings.TrimSpace(a[0])})
			default:
				c.result = append(c.result, &ResultColumn{ColumnName: strings.TrimSpace(col)})
			}

		}

		return nil
	}
}

// SetSelectOrderColumn 设置查询排序列
func SetSelectOrderColumn(col string, sort ...string) SelectOption {
	return func(c *SelectConfig) error {
		if len(col) > 0 {
			switch {
			case len(sort) > 1:
				return fmt.Errorf("only one sort column can be set for column %v", col)
			case len(sort) == 1:
				c.order = append(c.order, &OrderColumn{ColumnName: strings.TrimSpace(col), Order: strings.TrimSpace(sort[0])})
			default:
				c.order = append(c.order, &OrderColumn{ColumnName: strings.TrimSpace(col)})
			}
		}
		return nil
	}
}

// SetSelectPage 是否开启分页
func SetSelectPage() SelectOption {
	return func(c *SelectConfig) error {
		c.Page = true
		return nil
	}
}

// SetCountWhereColumn 设置count where
func SetCountWhereColumn(Operator, ColumnName, Expression string) CountOption {
	return func(c *CountConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 || len(Expression) == 0 || strings.EqualFold(strings.ToLower(Expression), "in") {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: Expression})
		return nil
	}
}

// SetUpdateSet 设置update set 列
func SetUpdateSet(s string) UpdateOption {
	return func(c *UpdateConfig) error {
		if len(s) == 0 {
			return fmt.Errorf("incorrect sql expression")
		}
		c.setColumnWithUpdate = append(c.setColumnWithUpdate, strings.TrimSpace(s))
		return nil
	}
}

func SetUpdateWhereColumn(Operator, ColumnName, Expression string) UpdateOption {
	return func(c *UpdateConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 || len(Expression) == 0 || strings.EqualFold(strings.ToLower(Expression), "in") {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: Expression})
		return nil
	}
}

func SetDeleteWhereColumn(Operator, ColumnName, Expression string) DeleteOption {
	return func(c *DeleteConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 || len(Expression) == 0 || strings.EqualFold(strings.ToLower(Expression), "in") {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: Expression})
		return nil
	}
}

func SetDeleteWhereColumnMustValue(Operator, ColumnName string, count int) DeleteOption {
	return func(c *DeleteConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: "in", mustValueCount: count})
		return nil
	}
}

func SetSelectWhereColumnMustValue(Operator, ColumnName string, count int) SelectOption {
	return func(c *SelectConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: "in", mustValueCount: count})
		return nil
	}
}

func SetUpdateWhereColumnMustValue(Operator, ColumnName string, count int) UpdateOption {
	return func(c *UpdateConfig) error {
		if len(Operator) == 0 || len(ColumnName) == 0 {
			return fmt.Errorf("incorrect sql expression")
		}
		c.predicate = append(c.predicate, &PredicateColumn{Operator: Operator, ColumnName: ColumnName, Expression: "in", mustValueCount: count})
		return nil
	}
}

// SetInsertColumn 设置insert 列
func SetInsertColumn(s string) InsertOption {
	return func(c *InsertConfig) error {
		if len(s) == 0 {
			return fmt.Errorf("incorrect sql expression")
		}
		c.columnWithInsert = append(c.columnWithInsert, strings.TrimSpace(s))
		return nil
	}
}
