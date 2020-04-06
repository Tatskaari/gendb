package main

import (
	"github.com/stretchr/testify/require"
	"github.com/tatskaari/gendb/recgen/template"
	"testing"
)

func TestParser(t *testing.T) {
	r := require.New(t)

	modelCode := `
		package model

		type FooRecord struct {
			ID string` + "`db:\"id\"`" + `
			BarID string` + "`db:\"bar_id\"`" + `
			Name string` + "`db:\"name\"`" + `
		}
	`

	pkg, cols, err := GetPackageAndStructFieldsFromFile(modelCode, "FooRecord", "foo")
	r.NoError(err)
	r.Equal("model", pkg)
	r.Len(cols, 3)

	r.Equal(template.Column{ColName:"foo.id", VariableName:"FooRecordID"}, cols[0])
	r.Equal(template.Column{ColName:"foo.bar_id", VariableName:"FooRecordBarID"}, cols[1])
	r.Equal(template.Column{ColName:"foo.name", VariableName:"FooRecordName"}, cols[2])
}
