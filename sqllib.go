package sqlbuilder

import "fmt"

const (
	setColumn     = "setColumnWithUpdate" // 用于update set、insert
	valueColumn   = "insertValues"        // 用于insert value
	result        = "ResultColumn"        // 用于select 返回列
	where         = "WhereColumn"         // 用于select、update、delete set
	whereNullCol  = "NullColPredicate"    // 用于select where null
	unknownWhere  = "UnknownWhere"        // 用于未知sql语法
	sort          = "SortColumn"          // 用于	select排序
	internalTable = "InternalTable"
)

var (
	InternalSelect     = "InternalSelect"
	InternalSelectText = fmt.Sprintf("select %s from %s where %s %s %s %s", result, internalTable, where, whereNullCol, unknownWhere, sort)
	InternalUpdate     = "InternalUpdate"
	InternalUpdateText = fmt.Sprintf("update %s set %s where %s", internalTable, setColumn, where)
	InternalDelete     = "InternalDelete"
	InternalDeleteText = fmt.Sprintf("delete from %s where %s", internalTable, where)
	InternalInsert     = "InternalInsert"
	InternalInsertText = fmt.Sprintf("insert into %s (%s) values(%s)", internalTable, setColumn, valueColumn)
)
