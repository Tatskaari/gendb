package sqlizer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tatskaari/gendb/builder"
)

func Select(sb *builder.SelectBuilder) (string, []interface{}) {
	if sb.Columns == nil {
		panic(errors.New("select builder was missing select columns"))
	}
	sql, args := selectCols(sb.Columns)

	if sb.FromTable != nil {
		sql = sql + " FROM " + sb.FromTable.Name
	}
	if sb.JoinBuilders != nil {
		joinClauses, joinArgs := joins(sb.JoinBuilders)
		sql = sql + " " + joinClauses
		args = combineArgs(args, joinArgs)
	}
	if sb.WhereBuilder != nil {
		whereClause, whereArgs := expr(sb.WhereBuilder.Expr)
		sql = sql + " WHERE " + whereClause
		args = combineArgs(args, whereArgs)
	}

	return sql, args
}

func selectCols(cols []builder.Expr) (string, []interface{}) {
	colsSql := make([]string, len(cols))
	var args []interface{}
	for i, col := range cols {
		colSql, colArgs := expr(col)
		colsSql[i] = colSql
		args = combineArgs(args, colArgs)
	}
	return "SELECT " + strings.Join(colsSql, ", "), args
}

func joins(joins []*builder.JoinBuilder) (string, []interface{}) {
	var args []interface{}
	joinSql := make([]string, len(joins))

	for i, j := range joins {
		sql, joinArgs := join(j)
		joinSql[i] = sql
		args = combineArgs(args, joinArgs)
	}

	return strings.Join(joinSql, " "), args
}

func join(join *builder.JoinBuilder) (string, []interface{}) {
	onCond, onArgs := expr(join.OnCondition.Expr)
	return fmt.Sprintf("JOIN %s ON %s", join.Table, onCond), onArgs
}
