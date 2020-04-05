package builder_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/tatskaari/gendb/builder"
	"testing"
)

type builderSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(builderSuite))
}

func (s *builderSuite) TestSelectFrom() {
	sb := builder.Select("name").From("foo")

	s.Equal("foo", sb.FromTable.Name)
	s.Require().Len(sb.Columns, 1, "incorrect number of columns in select")
	s.Require().IsType(&builder.IdentExpression{}, sb.Columns[0])
	s.Equal("name", sb.Columns[0].(*builder.IdentExpression).Name)
}

func (s *builderSuite) TestJoin() {
	jbs := builder.Select("name").
		From("foo").
		Join("bar").On(builder.Eq("foo.bar_id", "bar.id")).And(builder.Eq("foo.name", "bar.name")).
		Join("baz").On(builder.Eq("bar.baz_id", "baz.id")).
		JoinBuilders

	s.Require().Len(jbs, 2)

	s.Equal("bar", jbs[0].Table)
	s.Equal(&builder.BinOpExpr{
		LHS: &builder.BinOpExpr{
			LHS:    &builder.IdentExpression{Name: "foo.bar_id"},
			RHS:    &builder.IdentExpression{Name: "bar.id"},
			Symbol: "=",
		},
		RHS:    &builder.BinOpExpr{
			LHS:    &builder.IdentExpression{Name: "foo.name"},
			RHS:    &builder.IdentExpression{Name: "bar.name"},
			Symbol: "=",
		},
		Symbol: "AND",
	}, jbs[0].OnCondition.Expr)

	s.Equal("baz", jbs[1].Table)
	s.Equal(&builder.BinOpExpr{
		LHS:    &builder.IdentExpression{Name: "bar.baz_id"},
		RHS:    &builder.IdentExpression{Name: "baz.id"},
		Symbol: "=",
	}, jbs[1].OnCondition.Expr)
}

func (s *builderSuite) TestWhere() {
	wb := builder.Select().From("foo").
		Where(builder.Eq("name", builder.Bind("some_name"))).
		Or(builder.Col("active")).
		WhereBuilder

	s.Equal(&builder.BinOpExpr{
		LHS:    &builder.BinOpExpr{
			LHS:    &builder.IdentExpression{Name: "name"},
			RHS:    &builder.BoundValueExpr{Value: "some_name"},
			Symbol: "=",
		},
		RHS:    &builder.IdentExpression{Name: "active"},
		Symbol: "OR",
	}, wb.WhereBuilder.Expr)
}

func (s *builderSuite) TestToExpression() {
	s.Equal(&builder.IdentExpression{Name: "test"}, builder.ToExpression("test"))
	s.Equal(&builder.BoundValueExpr{Value: 5}, builder.ToExpression(5))
	s.Equal(&builder.IdentExpression{Name: "test"}, builder.Col("test"))
}

func (s *builderSuite) TestInsertBuilder() {
	ib := builder.Insert("foo").
		Values(map[string]interface{}{
			"id":     1234,
			"name":   builder.Bind("some name"),
			"bar_id": 2345,
		}).
		Values(map[string]interface{}{
			"id":     4321,
			"name":   builder.Bind("some other name"),
			"bar_id": 5432,
		})

	s.Equal(builder.Bind(1234), ib.ValueRows[0][ib.ColumnsOrder["id"]])
	s.Equal(builder.Bind(4321), ib.ValueRows[1][ib.ColumnsOrder["id"]])

	s.Equal(builder.Bind("some name"), ib.ValueRows[0][ib.ColumnsOrder["name"]])
	s.Equal(builder.Bind("some other name"), ib.ValueRows[1][ib.ColumnsOrder["name"]])

	s.Equal(builder.Bind(2345), ib.ValueRows[0][ib.ColumnsOrder["bar_id"]])
	s.Equal(builder.Bind(5432), ib.ValueRows[1][ib.ColumnsOrder["bar_id"]])
}

func (s *builderSuite) TestUpdateBuilder() {
	ub := builder.Update("foo").
		Set("bar_id", 1234).
		Set("name", builder.Bind("some_other_name")).
		Where(builder.Eq("id", 4321)).
		And(builder.Eq("name", builder.Bind("old_name")))

	s.Equal("foo", ub.Table)

	s.Require().Len(ub.Sets, 2)

	s.Equal(ub.Sets[0].Column.Name, "bar_id")
	s.Equal(ub.Sets[0].Value, builder.Bind(1234))

	s.Equal(ub.Sets[1].Column.Name, "name")
	s.Equal(ub.Sets[1].Value, builder.Bind("some_other_name"))

	s.Equal(builder.And(
		builder.Eq(builder.Col("id"), builder.Bind(4321)),
		builder.Eq(builder.Col("name"), builder.Bind("old_name")),
	), ub.WhereCondition.Expr)
}