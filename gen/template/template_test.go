package template_test

import (
	"github.com/stretchr/testify/suite"
	template "github.com/tatskaari/gendb/gen/template"
	"regexp"
	"testing"
)

type tempalteSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(tempalteSuite))
}

func (s *tempalteSuite) TestGenerate() {
	code, err := template.Generate("foo", []string{"id", "name", "bar_id"})
	s.Require().NoError(err)

	const expectedCode = `
	package foo
	
	import (
		"github.com/tatskaari/gendb/builder"
		"github.com/tatskaari/gendb/gen"
	)

	const TableName = "foo"

	var (
		Id = builder.Col("id")
		Name = builder.Col("name")
        BarId = builder.Col("bar_id")
		AllColumns = []string{
			"id",
			"name",
			"bar_id",
		}
	)

	func SelectFrom() *builder.SelectBuilder {return gen.SelectFrom(TableName, AllColumns)}
`
	s.Equal(normaliseWhitespace(expectedCode), normaliseWhitespace(code))
}

func normaliseWhitespace(text string) string {
	var re = regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(text, ` `)
}