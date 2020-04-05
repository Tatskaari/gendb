package sqlizer

import (
	"fmt"
	"github.com/tatskaari/gendb/builder"
)

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
