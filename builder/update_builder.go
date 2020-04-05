package builder

type UpdateBuilder struct {
	Table string
	Sets []*Set
	WhereCondition *UpdateExprBuilder
}

type Set struct {
	Column *IdentExpression
	Value Expr
}

func Update(table string) *UpdateBuilder {
	return &UpdateBuilder{
		Table: table,
	}
}

func (ub *UpdateBuilder) Set(column string, value interface{}) *UpdateBuilder {
	ub.Sets = append(ub.Sets, &Set{
		Column: Col(column),
		Value:  ToExpression(value),
	})
	return ub
}

func (ub *UpdateBuilder) Where(expr Expr) *UpdateExprBuilder {
	uxb := &UpdateExprBuilder{
		UpdateBuilder: ub,
		Expr:          expr,
	}

	ub.WhereCondition = uxb
	return uxb
}


type UpdateExprBuilder struct {
	*UpdateBuilder
	Expr Expr
}

func (eb *UpdateExprBuilder) And(expr Expr) *UpdateExprBuilder {
	eb.Expr = &AndExpr{
		LHS: eb.Expr,
		RHS: expr,
	}
	return eb
}

func (eb *UpdateExprBuilder) Or(expr Expr) *UpdateExprBuilder {
	eb.Expr = &OrExpr{
		LHS: eb.Expr,
		RHS: expr,
	}
	return eb
}