package builder

type SelectBuilder struct {
	Columns      []Expr
	FromTable  *IdentExpression
	JoinBuilders []*JoinBuilder
	WhereBuilder *SelectExprBuilder
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.FromTable = Col(table)
	return sb
}
func Select(columns ...string) *SelectBuilder {
	exprs := make([]Expr, len(columns))
	for i, colName := range columns {
		exprs[i] = &IdentExpression{Name: colName}
	}

	return &SelectBuilder{
		Columns: exprs,
	}
}

func (sb *SelectBuilder) Join(table string) *JoinBuilder {
	return &JoinBuilder{
		Parent:       sb,
		Table:        table,
	}
}

type JoinBuilder struct {
	Parent      *SelectBuilder
	Table       string
	OnCondition *SelectExprBuilder
}

func (jb *JoinBuilder) On(condition Expr) *SelectExprBuilder {
	jb.OnCondition = &SelectExprBuilder{
		SelectBuilder:  jb.Parent,
		Expr:         	condition,
	}
	jb.Parent.JoinBuilders = append(jb.Parent.JoinBuilders, jb)
	return jb.OnCondition
}


func (sb *SelectBuilder) Where(expr Expr) *SelectExprBuilder {
	sb.WhereBuilder = &SelectExprBuilder{
		SelectBuilder: sb,
		Expr:         	expr,
	}
	return sb.WhereBuilder
}

type SelectExprBuilder struct {
	*SelectBuilder
	Expr Expr
}

func (eb *SelectExprBuilder) And(expr Expr) *SelectExprBuilder {
	eb.Expr = &AndExpr{
		LHS: eb.Expr,
		RHS: expr,
	}

	return eb
}

func (eb *SelectExprBuilder) Or(expr Expr) *SelectExprBuilder {
	eb.Expr = &OrExpr{
		LHS: eb.Expr,
		RHS: expr,
	}

	return eb
}