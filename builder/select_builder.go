package builder

type selectExecutor interface {
	Query(sb *SelectBuilder, dest interface{}) error
	Get(sb *SelectBuilder, dest interface{}) error
}

type SelectBuilder struct {
	Columns      []Expr
	FromTable  *IdentExpression
	JoinBuilders []*JoinBuilder
	WhereBuilder *SelectExprBuilder

	Executor selectExecutor
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.FromTable = Col(table)
	return sb
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

func (eb *SelectExprBuilder) And(expr interface{}) *SelectExprBuilder {
	eb.Expr = And(eb.Expr, expr)
	return eb
}

func (eb *SelectExprBuilder) Or(expr interface{}) *SelectExprBuilder {
	eb.Expr = Or(eb.Expr, expr)
	return eb
}

func (sb *SelectBuilder) Query(dest interface{}) error {
	return sb.Executor.Query(sb, dest)
}

func (sb *SelectBuilder) Get(dest interface{}) error {
	return sb.Executor.Get(sb, dest)
}