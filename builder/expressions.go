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

type BinOpExpr struct {
	LHS Expr
	RHS Expr
	Symbol string
}

func (*BinOpExpr) isExpr() {}

func Eq(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: "=",
	}
}

func NotEq(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: "!=",
	}
}

func LT(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: "<",
	}
}

func LTE(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: "<=",
	}
}

func GT(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: ">",
	}
}

func GTE(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: ">=",
	}
}

func And(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: "AND",
	}
}

func Or(lhs interface{}, rhs interface{}) *BinOpExpr {
	return &BinOpExpr{
		LHS: ToExpression(lhs),
		RHS: ToExpression(rhs),
		Symbol: "OR",
	}
}

type UnaryOpExpr struct {
	Expr Expr
	Symbol string
}

func Not(expr interface{}) *UnaryOpExpr {
	return &UnaryOpExpr{
		Expr:   ToExpression(expr),
		Symbol: "NOT",
	}
}

func (*UnaryOpExpr) isExpr() {}

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
