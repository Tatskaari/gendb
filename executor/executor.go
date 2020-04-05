package executor

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/tatskaari/gendb/builder"
)

type QueryRunner interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Rebind(query string) string
}

type Sqlizer interface {
	Select(selectBuilder *builder.SelectBuilder) (sql string, args []interface{})
	Insert(interfaceBuilder *builder.InsertBuilder) (sql string, args []interface{})
	Update(updateBuilder *builder.UpdateBuilder) (sql string, args []interface{})
}

type Executor struct {
	qr QueryRunner
	sqlizer Sqlizer
}

func New(qr QueryRunner, sqlizer Sqlizer) *Executor {
	return &Executor{
		qr:      qr,
		sqlizer: sqlizer,
	}
}

func (e *Executor) Insert() *builder.InsertBuilder {
	return &builder.InsertBuilder{
		Executor: e,
	}
}

func (e *Executor) Update(table string) *builder.UpdateBuilder {
	ub := &builder.UpdateBuilder{
		Executor: e,
	}
	return ub.Update(table)
}

func (e *Executor) Select(columns ...string) *builder.SelectBuilder {
	sb := &builder.SelectBuilder{
		Executor: e,
	}
	return sb.Select(columns...)
}

func (e *Executor) ExecInsert(ib *builder.InsertBuilder) (sql.Result, error) {
	s, args := e.sqlizer.Insert(ib)
	return e.qr.Exec(e.qr.Rebind(s), args...)
}

func (e *Executor) ExecUpdate(ib *builder.UpdateBuilder) (sql.Result, error) {
	s, args := e.sqlizer.Update(ib)
	return e.qr.Exec(e.qr.Rebind(s), args...)
}

func (e *Executor) Query(sb *builder.SelectBuilder, dest interface{}) error {
	s, args := e.sqlizer.Select(sb)
	rows, err := e.qr.Query(e.qr.Rebind(s), args...)
	if err != nil {
		return err
	}

	return sqlx.StructScan(rows, dest)
}