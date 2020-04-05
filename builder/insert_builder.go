package builder

import (
	"database/sql"
)

type insertExecutor interface {
	ExecInsert(builder *InsertBuilder) (sql.Result, error)
}

type InsertBuilder struct {
	IntoTable string
	ColumnsOrder map[string]int
	ValueRows		 [][]Expr

	Executor insertExecutor
}

func (ib *InsertBuilder) Into(into string) *InsertBuilder {
	ib.IntoTable = into
	return ib
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

func (ib *InsertBuilder) Exec() (sql.Result, error) {
	return ib.Executor.ExecInsert(ib)
}