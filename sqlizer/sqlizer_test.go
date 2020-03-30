package sqlizer_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/tatskaari/gendb/builder"
	"github.com/tatskaari/gendb/sqlizer"
	"testing"
)

type sqlizerSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(sqlizerSuite))
}

func (s *sqlizerSuite) TestSelectFrom() {
	sb := builder.Select("name").From("foo")
	sql, args := sqlizer.Sqlize(sb)

	s.Equal("SELECT name FROM foo", sql)
	s.Nil(args)
}

func (s *sqlizerSuite) TestJoin() {
	sb := builder.Select("foo.name").
		From("foo").
		Join("bar").On(builder.Eq("foo.bar_id", "bar.id")).And(builder.Col("active")).
		Join("baz").On(builder.Eq("bar.baz_id", "baz.id")).Or(builder.Col("active"))

	sql, args := sqlizer.Sqlize(sb.SelectBuilder)

	s.Equal("SELECT foo.name FROM foo JOIN bar ON foo.bar_id = bar.id AND active JOIN baz ON bar.baz_id = baz.id OR active", sql)
	s.Nil(args)
}

func (s *sqlizerSuite) TestWhere() {
	sb := builder.Select("foo.id").
		From("foo").
		Where(builder.Eq("name", builder.Bind("name"))).And(builder.Col("active"))

	sql, args := sqlizer.Sqlize(sb.SelectBuilder)

	s.Equal("SELECT foo.id FROM foo WHERE name = ? AND active", sql)
	s.Require().Len(args, 1)
	s.Require().IsType("", args[0])
	s.Equal("name", args[0])
}

