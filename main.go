package main

import (
	"astGenerator/internal/ast_parser"
	"astGenerator/internal/generator"
	"astGenerator/internal/template"
	"fmt"
	"os"
)

func main() {
	parser := ast_parser.New()
	data, err := parser.ParseFile("internal/generator/fixtures/06_struct.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gen := generator.New(template.MsgpackEncDecTmpl)
	res, err := gen.Generate(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Create("internal/generator/fixtures/06_struct_enc_exp.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err = f.WriteString(res); err != nil {
		panic(err)
	}
}
