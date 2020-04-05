package sqlizer

import (
	"github.com/tatskaari/gendb/builder"
	"strings"
)

func Update(ub *builder.UpdateBuilder) (string, []interface{}) {
	setSQl, args := sets(ub.Sets)
	sql := "UPDATE " + ub.Table + " " + setSQl

	if ub.WhereCondition != nil {
		whereClause, whereArgs := expr(ub.WhereCondition.Expr)
		sql = sql + " WHERE " + whereClause
		args = combineArgs(args, whereArgs)
	}
	return sql, args
}

func sets(sets []*builder.Set) (string, []interface{}) {
	var args []interface{}
	setSqlFragments := make([]string, len(sets))

	for i, s := range sets {
		sql, a := set(s)
		args = combineArgs(args, a)
		setSqlFragments[i] = sql
	}
	return "SET " + strings.Join(setSqlFragments, ", "), args
}

func set(set *builder.Set) (string, []interface{}) {
	valueSql, args := expr(set.Value)
	return set.Column.Name + " = " + valueSql, args
}