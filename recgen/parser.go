package main

import (
	"fmt"
	"github.com/tatskaari/gendb/recgen/template"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

func GetPackageAndStructFieldsFromFile(file string, structName string, tableName string) (string, []template.Column, error) {

	fileAST, err := parser.ParseFile(token.NewFileSet(), "", file, parser.AllErrors)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get package: %w", err)
	}

	ast.FilterFile(fileAST, func(s string) bool {
		return s == structName
	})

	recordStruct := getStructDefinition(fileAST, structName)
	cols, err := getColumns(structName, tableName, recordStruct)

	return getPackageName(fileAST), cols, err
}

func getColumns(structName string, tableName string, recordStruct *ast.StructType) ([]template.Column, error) {
	var columns []template.Column
	for _, field := range recordStruct.Fields.List {
		tagString := field.Tag.Value

		tag := reflect.StructTag(tagString[1:(len(tagString) - 1)])
		dbName := tag.Get("db")

		if dbName == "" {
			return nil, fmt.Errorf("failed to get db name from struct tag: %s", string(tag))
		}

		for _, name := range field.Names {
			columns = append(columns, template.Column{
				VariableName: structName + name.Name,
				ColName:      tableName + "." + dbName,
			})
		}
	}
	return columns, nil
}

func getPackageName(file *ast.File) string {
	return file.Name.Name
}

func getStructDefinition(file *ast.File, name string) *ast.StructType {
	for _, dec := range file.Decls {
		if genDec, ok := dec.(*ast.GenDecl); ok {
			for _, spec := range genDec.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						if typeSpec.Name.Name == name {
							return structType
						}
					}
				}
			}
		}
	}
	return nil
}
