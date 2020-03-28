package builder

type SelectBuilder struct {
	Columns      []Expr
	From  string
	JoinBuilders []*JoinBuilder
	WhereBuilder *ExprBuilder
}

func From(table string) *SelectBuilder {
	return &SelectBuilder{
		From: table,
	}
}
func (sb *SelectBuilder) Select(columns ...string) *SelectBuilder {
	exprs := make([]Expr, len(columns))
	for i, colName := range columns {
		exprs[i] = &IdentExpression{Name: colName}
	}

	sb.Columns = exprs
	return sb
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