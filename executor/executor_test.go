package executor_test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tatskaari/gendb/builder"
	"github.com/tatskaari/gendb/executor"
	"github.com/tatskaari/gendb/sqlizer"
	"testing"
)

type fooRecord struct {
	ID int `db:"id"`
	BarID int `db:"bar_id"`
	Name string `db:"name"`
}

type executorTestSuite struct {
	suite.Suite
	con *sql.DB
}

func TestGenDB(t *testing.T) {
	suite.Run(t, new(executorTestSuite))
}

func (s *executorTestSuite) SetupTest() {
	con, err := sql.Open("sqlite3", ":memory:")
	s.NoError(err, "failed to start database")
	s.con = con

	_, err = con.Exec("CREATE TABLE foo (id INTEGER, bar_id INTEGER, name TEXT)")
	s.NoError(err)

	_, err = con.Exec("CREATE TABLE bar (id INTEGER, name TEXT)")
	s.NoError(err)

}

func (s *executorTestSuite) TestExecute() {
	ex := executor.New(s.con, new(sqlizer.StandardSqlizer))

	expectedFooRow := fooRecord{
		ID:    1234,
		BarID: 4567,
		Name:  "some_name",
	}

	_, err := ex.Insert().Into("foo").Values(map[string]interface{} {
		"id": expectedFooRow.ID,
		"bar_id": expectedFooRow.BarID,
		"name": builder.Bind(expectedFooRow.Name),
	}).Exec()
	s.Require().NoError(err)

	var fooRows []fooRecord
	err = ex.Select("id", "name", "bar_id").
		From("foo").
		Where(builder.Eq("id", expectedFooRow.ID)).
		Query(&fooRows)

	s.Require().Len(fooRows, 1)
	s.Equal(expectedFooRow, fooRows[0])

	_, err = ex.Update("foo").
		Set("name", builder.Bind("some other name")).
		Where(builder.Eq("id", expectedFooRow.ID)).
		Exec()

	s.Require().NoError(err)

	fooRows = nil
	err = ex.Select("id", "name", "bar_id").
		From("foo").
		Where(builder.Eq("id", expectedFooRow.ID)).
		Query(&fooRows)

	s.Require().Len(fooRows, 1)
	s.Equal("some other name", fooRows[0].Name)
}
