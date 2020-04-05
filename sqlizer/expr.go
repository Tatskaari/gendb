package sqlizer

import (
	"fmt"
	"github.com/tatskaari/gendb/builder"
)

func expr(expr builder.Expr) (string, []interface{}) {
	switch expr.(type) {
	case *builder.IdentExpression:
		return expr.(*builder.IdentExpression).Name, nil
	case *builder.BinOpExpr:
		return binOpExpr(expr.(*builder.BinOpExpr))
	case *builder.BoundValueExpr:
		return boundValExpr(expr.(*builder.BoundValueExpr))
	default:
		panic(fmt.Errorf("unknown expression type: %#v", expr))
	}
}

func binOpExpr(binOpExpr *builder.BinOpExpr) (string, []interface{}) {
	lhsSql, lhsArgs := expr(binOpExpr.LHS)
	rhsSql, rhsArgs := expr(binOpExpr.RHS)

	return fmt.Sprintf("%s %s %s", lhsSql, binOpExpr.Symbol, rhsSql), combineArgs(lhsArgs, rhsArgs)
}

func boundValExpr(expr *builder.BoundValueExpr) (string, []interface{}) {
	return "?", []interface{} {expr.Value}
}