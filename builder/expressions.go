package builder

type Expr interface {
	isExpr()
}


type IdentExpression struct {
	Name string
}

func (*IdentExpression) isExpr() {}

func Col(name string) *IdentExpression {
	return &IdentExpression{Name: name}
}

type EqExpression struct {
	LHS Expr
	RHS Expr
}

func (*EqExpression) isExpr() {}

func ColEq(lhs string, rhs string) *EqExpression {
	return &EqExpression{
		LHS: &IdentExpression{lhs},
		RHS: &IdentExpression{rhs},
	}
}

func Eq(lhs Expr, rhs Expr) *EqExpression {
	return &EqExpression{
		LHS: lhs,
		RHS: rhs,
	}
}

type BoundValueExpr struct {
	Value interface{}
}

func (*BoundValueExpr) isExpr() {}

func Bind(value interface{}) *BoundValueExpr {
	return &BoundValueExpr{
		Value: value,
	}
}

type AndExpr struct {
	LHS Expr
	RHS Expr
}

func (*AndExpr) isExpr() {}

type OrExpr struct {
	LHS Expr
	RHS Expr
}

func (*OrExpr) isExpr() {}