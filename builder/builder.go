package builder

type SelectBuilder struct {
	Columns      []Expr
	FromBuilder  *FromBuilder
	JoinBuilders []*JoinBuilder
	WhereBuilder *ExprBuilder
}


type FromBuilder struct {
	Table string
}

func From(table string) *SelectBuilder {
	return &SelectBuilder{
		FromBuilder: &FromBuilder{
			Table: table,
		},
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

type JoinBuilder struct {
	Parent      *SelectBuilder
	Table       string
	OnCondition *ExprBuilder
}

func (sb *SelectBuilder) Join(table string) *JoinBuilder {
	return &JoinBuilder{
		Parent:       sb,
		Table:        table,
	}
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