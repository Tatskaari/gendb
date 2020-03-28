package sqlizer

import (
	"errors"
	"fmt"
	"gendb/builder"
	"strings"
)

func Sqlize(sb *builder.SelectBuilder) (string, []interface{}) {
	if sb.Columns == nil {
		panic(errors.New("select builder was missing select columns"))
	}
	sql, args := selectCols(sb.Columns)

	if sb.From != "" {
		sql = sql + " FROM " + sb.From
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

func combineArgs(args []interface{}, newArgs []interface{}) []interface{} {
	combinedArgs := args
	for _, arg := range newArgs {
		combinedArgs = append(combinedArgs, arg)
	}
	return combinedArgs
}

func expr(expr builder.Expr) (string, []interface{}) {
	switch expr.(type) {
	case *builder.IdentExpression:
		return expr.(*builder.IdentExpression).Name, nil
	case *builder.EqExpression:
		return eqExpr(expr.(*builder.EqExpression))
	case *builder.AndExpr:
		return andExpr(expr.(*builder.AndExpr))
	case *builder.OrExpr:
		return orExpr(expr.(*builder.OrExpr))
	case *builder.BoundValueExpr:
		return boundValExpr(expr.(*builder.BoundValueExpr))
	default:
		panic(fmt.Errorf("unknown expression type: %#v", expr))
	}
}

// TODO: Make a "binary expression" interface that has a LHS() and RHS() method do reduce duplication here
func eqExpr(eqExpr *builder.EqExpression) (string, []interface{}) {
	lhsSql, lhsArgs := expr(eqExpr.LHS)
	rhsSql, rhsArgs := expr(eqExpr.RHS)

	return fmt.Sprintf("%s = %s", lhsSql, rhsSql), combineArgs(lhsArgs, rhsArgs)
}

func andExpr(eqExpr *builder.AndExpr) (string, []interface{}) {
	lhsSql, lhsArgs := expr(eqExpr.LHS)
	rhsSql, rhsArgs := expr(eqExpr.RHS)

	return fmt.Sprintf("%s AND %s", lhsSql, rhsSql), combineArgs(lhsArgs, rhsArgs)
}

func orExpr(eqExpr *builder.OrExpr) (string, []interface{}) {
	lhsSql, lhsArgs := expr(eqExpr.LHS)
	rhsSql, rhsArgs := expr(eqExpr.RHS)

	return fmt.Sprintf("%s OR %s", lhsSql, rhsSql), combineArgs(lhsArgs, rhsArgs)
}

func boundValExpr(expr *builder.BoundValueExpr) (string, []interface{}) {
	return "?", []interface{} {expr.Value}
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
