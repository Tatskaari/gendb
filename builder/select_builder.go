package builder

type SelectBuilder struct {
	Columns      []Expr
	FromTable  *IdentExpression
	JoinBuilders []*JoinBuilder
	WhereBuilder *ExprBuilder
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
	OnCondition *ExprBuilder
}



func (jb *JoinBuilder) On(condition Expr) *ExprBuilder {
	jb.OnCondition = &ExprBuilder{
		SelectBuilder:  jb.Parent,
		Expr:         	condition,
	}
	jb.Parent.JoinBuilders = append(jb.Parent.JoinBuilders, jb)
	return jb.OnCondition
}


func (sb *SelectBuilder) Where(expr Expr) *ExprBuilder {
	sb.WhereBuilder = &ExprBuilder{
		SelectBuilder: sb,
		Expr:         	expr,
	}
	return sb.WhereBuilder
}