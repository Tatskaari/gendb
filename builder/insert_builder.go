package builder

type InsertBuilder struct {
	Into string
	ColumnsOrder map[string]int
	ValueRows		 [][]Expr
}

func Insert(into string) *InsertBuilder {
	return &InsertBuilder{
		Into: into,
	}
}

func (ib *InsertBuilder) Values(values map[string]interface{}) *InsertBuilder {
	if ib.ColumnsOrder == nil {
		i := 0
		ib.ColumnsOrder = make(map[string]int, len(values))
		for k := range values {
			ib.ColumnsOrder[k] = i
			i++
		}
	}

	vs := make([]Expr, len(values))
	for k, v := range values {
		vs[ib.ColumnsOrder[k]] = ToExpression(v)
	}
	ib.ValueRows = append(ib.ValueRows, vs)
	return ib
}
