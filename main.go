package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {

	_, err := parser.ParseFile(token.NewFileSet(), "test.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	ts := &ast.TypeSpec{
		Name: &ast.Ident{Name: "Test"},
		Type:    &ast.StructType{
			Fields: new(ast.FieldList),
		},
	}
	var buff bytes.Buffer
	err = format.Node(&buff, token.NewFileSet(), ts)

	if err != nil {
		panic(err)
	}

	fmt.Printf(buff.String())
}
