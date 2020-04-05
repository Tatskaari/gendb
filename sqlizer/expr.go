package sqlizer

import (
	"fmt"
	"github.com/tatskaari/gendb/builder"
)

func Expr(expr builder.Expr) (string, []interface{}) {
	switch expr.(type) {
	case *builder.IdentExpression:
		return expr.(*builder.IdentExpression).Name, nil
	case *builder.BinOpExpr:
		return binOpExpr(expr.(*builder.BinOpExpr))
	case *builder.UnaryOpExpr:
		return unaryOpExpr(expr.(*builder.UnaryOpExpr))
	case *builder.BoundValueExpr:
		return boundValExpr(expr.(*builder.BoundValueExpr))
	default:
		panic(fmt.Errorf("unknown expression type: %#v", expr))
	}
}

func binOpExpr(binOpExpr *builder.BinOpExpr) (string, []interface{}) {
	lhsSql, lhsArgs := Expr(binOpExpr.LHS)
	rhsSql, rhsArgs := Expr(binOpExpr.RHS)

	return fmt.Sprintf("%s %s %s", lhsSql, binOpExpr.Symbol, rhsSql), combineArgs(lhsArgs, rhsArgs)
}

func unaryOpExpr(binOpExpr *builder.UnaryOpExpr) (string, []interface{}) {
	exprSql, args := Expr(binOpExpr.Expr)

	return fmt.Sprintf("%s %s", binOpExpr.Symbol, exprSql), args
}

func boundValExpr(expr *builder.BoundValueExpr) (string, []interface{}) {
	return "?", []interface{} {expr.Value}
}