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
	s.Require().Len( sb.Columns, 1, "incorrect number of columns in select")
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
	joinBuilder := jbs[0]

	s.Equal("bar", joinBuilder.Table)

	s.Require().IsType(&builder.AndExpr{}, jbs[0].OnCondition.Expr)
	andExpr := joinBuilder.OnCondition.Expr.(*builder.AndExpr)

	s.Require().IsType(&builder.EqExpression{}, andExpr.LHS)
	andLHS := andExpr.LHS.(*builder.EqExpression)

	s.Require().IsType(&builder.IdentExpression{}, andLHS.LHS)
	s.Equal("foo.bar_id", andLHS.LHS.(*builder.IdentExpression).Name)

	s.Require().IsType(&builder.IdentExpression{}, andLHS.RHS)
	s.Equal("bar.id", andLHS.RHS.(*builder.IdentExpression).Name)

	s.Require().IsType(&builder.EqExpression{}, andExpr.RHS)
	andRHS := andExpr.RHS.(*builder.EqExpression)

	s.Require().IsType(&builder.IdentExpression{}, andLHS.LHS)
	s.Equal("foo.name", andRHS.LHS.(*builder.IdentExpression).Name)

	s.Require().IsType(&builder.IdentExpression{}, andLHS.RHS)
	s.Equal("bar.name", andRHS.RHS.(*builder.IdentExpression).Name)

	joinBuilder = jbs[1]

	s.Equal("baz", joinBuilder.Table)

	s.Require().IsType(&builder.EqExpression{}, jbs[1].OnCondition.Expr)
	onCondition := joinBuilder.OnCondition.Expr.(*builder.EqExpression)

	s.Require().IsType(&builder.IdentExpression{}, onCondition.LHS)
	s.Equal("bar.baz_id", onCondition.LHS.(*builder.IdentExpression).Name)

	s.Require().IsType(&builder.IdentExpression{}, onCondition.RHS)
	s.Equal("baz.id", onCondition.RHS.(*builder.IdentExpression).Name)

}

func (s *builderSuite) TestWhere() {
	wb := builder.Select().From("foo").
		Where(builder.Eq("name", builder.Bind("some_name"))).
		Or(builder.Col("active")).
		WhereBuilder

	s.Require().IsType(&builder.OrExpr{}, wb.Expr)
	andExpr := wb.Expr.(*builder.OrExpr)

	s.Require().IsType(&builder.EqExpression{}, andExpr.LHS)
	s.Require().IsType(&builder.IdentExpression{}, andExpr.RHS)
}

func (s *builderSuite) TestToExpression(){
	s.Equal(&builder.IdentExpression{Name: "test"}, builder.ToExpression("test"))
	s.Equal(&builder.BoundValueExpr{Value: 5}, builder.ToExpression(5))
	s.Equal(&builder.IdentExpression{Name: "test"}, builder.Col("test"))
}