package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/tatskaari/gendb/recgen/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var opts struct {
	TableName string `long:"table_name" short:"t"`
	RecordStruct string `long:"record_struct" short:"r"`
	File string `long:"file" short:"f"`
	Output string `long:"out" short:"o" default:""`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	if opts.Output == "" {
		opts.Output = fmt.Sprintf("%s_record_utils.go", opts.TableName)
	}

	file, err := ioutil.ReadFile(opts.File)
	if err != nil {
		panic(err)
	}

	packageName, colums, err := GetPackageAndStructFieldsFromFile(string(file), opts.RecordStruct, opts.TableName)
	if err != nil {
		panic(err)
	}

	code, err :=  template.Generate(template.Data{
		Package:   packageName,
		CMD:       strings.Join(os.Args, " "),
		TableName: opts.TableName,
		Columns:   colums,
	})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(opts.Output, []byte(code), 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "fmt", opts.Output)
	err = cmd.Start()
	if err != nil {
		panic(fmt.Sprintf("failed to format output: %v",  err))
	}
	_, err = cmd.Process.Wait()
	if err != nil {
		panic(fmt.Sprintf("failed to format output: %v",  err))
	}
}
