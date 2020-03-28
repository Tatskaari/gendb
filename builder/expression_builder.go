package builder

type ExprBuilder struct {
	*SelectBuilder
	Expr Expr
}

func (eb *ExprBuilder) And(expr Expr) *ExprBuilder {
	eb.Expr = &AndExpr{
		LHS: eb.Expr,
		RHS: expr,
	}

	return eb
}

func (eb *ExprBuilder) Or(expr Expr) *ExprBuilder {
	eb.Expr = &OrExpr{
		LHS: eb.Expr,
		RHS: expr,
	}

	return eb
}