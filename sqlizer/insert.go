package sqlizer

import (
	"github.com/tatskaari/gendb/builder"
	"strings"
)

func Insert(ib *builder.InsertBuilder) (string, []interface{}) {
	values, args := values(ib.ValueRows)
	return "INSERT INTO " + ib.Into + " (" + insertCols(ib.ColumnsOrder) + ") " + values, args
}

func insertCols(columnOrders map[string]int) string {
	columns := make([]string, len(columnOrders))
	for col, order := range columnOrders {
		columns[order] = col
	}
	return strings.Join(columns, ", ")
}

func values(rows [][]builder.Expr) (string, []interface{}) {
	var args []interface{}
	rowSql := make([]string, len(rows))
	for i, row := range rows {
		sql, a := valueRow(row)
		rowSql[i] = sql
		args = combineArgs(args, a)
	}

	return "VALUES " + strings.Join(rowSql, ", "), args
}

func valueRow(values []builder.Expr) (string, []interface{}) {
	var args []interface{}
	valueSql := make([]string, len(values))

	for i, value := range values {
		sql, a := expr(value)
		args = combineArgs(args, a)
		valueSql[i] = sql
	}

	return "(" + strings.Join(valueSql, ", ") + ")", args
}