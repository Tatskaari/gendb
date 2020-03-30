package template

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/tatskaari/gendb/builder"
	"strings"
)

const templateString = `
	package {{TableName}}
	
	import (
		"github.com/tatskaari/gendb/builder"
		"github.com/tatskaari/gendb/gen"
	)

	const TableName = "{{TableName}}"

	var (
		{{#Columns}}
		{{VariableName}} = builder.Col("{{ColName}}")
		{{/Columns}}
		AllColumns = []string{
			{{#Columns}}
			"{{ColName}}",
			{{/Columns}}
		}
	)
	
	func SelectFrom() *builder.SelectBuilder {return gen.SelectFrom(TableName, AllColumns)}
`

func SelectFrom(table string, columns []string) *builder.SelectBuilder {
	return builder.Select(columns...).From(table)
}

type data struct {
	TableName string
	Columns []column
}

type column struct {
	VariableName string
	ColName string
}

func Generate(tableName string, columnNames []string) (string, error) {
	template, err := mustache.ParseString(templateString)
	if err != nil {
		return "", fmt.Errorf("failed to parse table mustache template: %w", err)
	}

	d := &data{
		TableName: tableName,
	}

	for _, columnName := range columnNames {
		column := column{
			VariableName: formatColumnNameToVarName(columnName),
			ColName:      columnName,
		}
		d.Columns = append(d.Columns, column)
	}

	res, err := template.Render(d)
	if err != nil {
		return "", fmt.Errorf("failed to redner table mustache template, data: %v: %w", d, err)
	}
	return res, nil
}

func formatColumnNameToVarName(name string) string {
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}