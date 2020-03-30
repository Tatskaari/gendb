package builder

type Expr interface {
	isExpr()
}


type IdentExpression struct {
	Name string
}

func (*IdentExpression) isExpr() {}

func Col(name string) *IdentExpression {
	// TODO: sanitize the name of any injected SQL
	return &IdentExpression{Name: name}
}

type EqExpression struct {
	LHS Expr
	RHS Expr
}

func (*EqExpression) isExpr() {}

func Eq(lhs interface{}, rhs interface{}) *EqExpression {
	return &EqExpression{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
	}
}

func ToExpression(expr interface{}) Expr {
	if expr, ok := expr.(Expr); ok {
		return expr
	}
	if expr, ok := expr.(string); ok {
		return Col(expr)
	}
	return Bind(expr)
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