package sqlizer_test

import (
	"fmt"
	"github.com/tatskaari/gendb/builder"
	"github.com/tatskaari/gendb/sqlizer"
)

func (s *sqlizerSuite) TestBinaryExpressions() {
	testCases := []struct {
		expr   builder.Expr
		symbol string
		args   []interface{}
	} {
		{
			expr:   builder.Eq(123, 321),
			symbol: "=",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.NotEq(123, 321),
			symbol: "!=",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.LT(123, 321),
			symbol: "<",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.GT(123, 321),
			symbol: ">",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.LTE(123, 321),
			symbol: "<=",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.GTE(123, 321),
			symbol: ">=",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.And(123, 321),
			symbol: "AND",
			args:   []interface{} {123, 321},
		},
		{
			expr:   builder.Or(123, 321),
			symbol: "OR",
			args:   []interface{} {123, 321},
		},
	}

	for _, testCase := range testCases {
		s.Run(fmt.Sprintf("%s binary op", testCase.symbol), func() {
			sql, args := sqlizer.Expr(testCase.expr)
			s.Equal(fmt.Sprintf("? %s ?", testCase.symbol), sql)
			s.Equal(testCase.args, args)
		})
	}
}

func (s *sqlizerSuite) TestUnaryExpressions() {
	testCases := []struct {
		expr   builder.Expr
		symbol string
		args   []interface{}
	} {
		{
			expr:   builder.Not(true),
			symbol: "NOT",
			args:   []interface{} {true},
		},
	}

	for _, testCase := range testCases {
		s.Run(fmt.Sprintf("%s unary op", testCase.symbol), func() {
			sql, args := sqlizer.Expr(testCase.expr)
			s.Equal(fmt.Sprintf("%s ?", testCase.symbol), sql)
			s.Equal(testCase.args, args)
		})
	}
}