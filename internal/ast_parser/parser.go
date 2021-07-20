package ast_parser

import (
	"astGenerator/internal/template"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

const TagName = "msgpack_index"

var typeEncoders = map[string]string{
	"int":          "Int",
	"int8":         "Int8",
	"int16":        "Int16",
	"int32":        "Int32",
	"int64":        "Int64",
	"uint":         "Uint",
	"uint8":        "Uint8",
	"uint16":       "Uint16",
	"uint32":       "Uint32",
	"uint64":       "Uint64",
	"float32":      "Float32",
	"float64":      "Float64",
	"bool":         "Bool",
	"string":       "String",
	"CustomInt":    "Int",
	"CustomString": "String",
}

type AstParser struct {
}

func New() *AstParser {
	return &AstParser{}
}

func (a *AstParser) ParseFile(filename string) ([]*template.MsgpackEncDecTemplateData, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("error parse file %s: %w", filename, err)
	}

	res := make([]*template.MsgpackEncDecTemplateData, 0)

	for _, decl := range f.Decls {
		if d, ok := decl.(*ast.GenDecl); ok {
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					if tspec, ok := spec.(*ast.TypeSpec); ok {
						if t, ok := tspec.Type.(*ast.StructType); ok {
							templateData := &template.MsgpackEncDecTemplateData{
								Package:               f.Name.String(),
								Fields:                make([]template.Field, 0),
								StructName:            tspec.Name.Name,
								StructNameFirstLetter: strings.ToLower(string([]rune(tspec.Name.Name)[0])),
							}

							for _, field := range t.Fields.List {

								var fieldType string
								switch field.Type.(type) {
								case *ast.Ident:
									fieldType = field.Type.(*ast.Ident).Name

									templateData.Fields = append(
										templateData.Fields,
										template.Field{
											ConvType:    strings.ToLower(typeEncoders[fieldType]),
											MsgPackType: typeEncoders[fieldType],
											Name:        field.Names[0].Name,
											Index:       a.getMbIndex(field.Tag.Value),
										},
									)
								default:
									templateData.Fields = append(
										templateData.Fields,
										template.Field{
											ConvType:    "",
											MsgPackType: "",
											Name:        field.Names[0].Name,
											Index:       a.getMbIndex(field.Tag.Value),
										},
									)
								}
							}
							res = append(res, templateData)
						}
					}
				}
			}
		}
	}

	return res, nil
}

func (a *AstParser) getMbIndex(tag string) int64 {
	re := regexp.MustCompile(TagName + `:"([0-9]*)"`)
	res := re.FindAllStringSubmatch(tag, -1)
	if len(res) <= 0 {
		return -1
	}

	num, err := strconv.ParseInt(res[len(res)-1][1], 10, 64)
	if err != nil {
		return -1
	}

	return num
}
