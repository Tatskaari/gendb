package builder

import "database/sql"

type updateExecutor interface {
	ExecUpdate(builder *UpdateBuilder) (sql.Result, error)
}

type UpdateBuilder struct {
	Table string
	Sets []*Set
	WhereCondition *UpdateExprBuilder

	Executor updateExecutor
}

type Set struct {
	Column *IdentExpression
	Value Expr
}

func (ub *UpdateBuilder) Update(table string) *UpdateBuilder {
	ub.Table = table
	return ub
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

func (eb *UpdateExprBuilder) And(expr interface{}) *UpdateExprBuilder {
	eb.Expr = And(eb.Expr, expr)
	return eb
}

func (eb *UpdateExprBuilder) Or(expr interface{}) *UpdateExprBuilder {
	eb.Expr = Or(eb.Expr, expr)
	return eb
}

func (ub *UpdateBuilder)  Exec() (sql.Result, error) {
	return ub.Executor.ExecUpdate(ub)
}