package generator

import (
	"astGenerator/internal/template"
	"bytes"
	"fmt"
	"sort"
	template2 "text/template"
)

type Generator struct {
	template string
}

func New(template string) *Generator {
	return &Generator{
		template: template,
	}
}

func (g *Generator) Generate(data []*template.MsgpackEncDecTemplateData) (string, error) {
	if len(data) <= 0 {
		return "", nil
	}

	var bufStr []byte
	buf := bytes.NewBuffer(bufStr)

	res := `//CODE GENERATED AUTOMATICALLY
//THIS FILE SHOULD NOT BE EDITED BY HAND
package ` + data[0].Package + `

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)
`

	for _, d := range data {
		sort.SliceStable(
			d.Fields,
			func(i, j int) bool {
				return d.Fields[i].Index < d.Fields[j].Index
			},
		)

		d.FieldsCount = len(d.Fields)
		t := template2.Must(template2.New("template-1").Parse(g.template))

		buf.Reset()
		err := t.Execute(buf, d)
		if err != nil {
			return "", fmt.Errorf("error on execute template: %w", err)
		}

		res += buf.String()

	}

	return res, nil
}
