package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/tatskaari/gendb/executor"
	"github.com/tatskaari/gendb/sqlizer"
)

func InTransaction(con *sqlx.DB, txFunc func(dao *DAO) error) (err error) {
	tx, err := con.Beginx()
	if err != nil {
		return err
	}
	dao := New(executor.New(tx, new(sqlizer.StandardSqlizer)))
	err = txFunc(dao)

	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}