package generator

import (
	"astGenerator/internal/ast_parser"
	"astGenerator/internal/template"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

func Test_Generate(t *testing.T) {
	var re = regexp.MustCompile(`[[:space:]]`)

	for i := 1; i <= 6; i++ {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			sourceFilePath := fmt.Sprintf("fixtures/0%d_struct.go", i)
			expectedFilePath := fmt.Sprintf("fixtures/0%d_struct_enc_exp.go", i)

			parser := ast_parser.New()
			data, err := parser.ParseFile(sourceFilePath)
			assert.NoError(t, err)

			gen := New(template.MsgpackEncDecTmpl)
			res, err := gen.Generate(data)
			assert.NoError(t, err)

			actualBytes := re.ReplaceAllString(string(res), "")

			fileEncExp, err := ioutil.ReadFile(expectedFilePath)
			assert.NoError(t, err)
			expectedBytes := re.ReplaceAllString(string(fileEncExp), "")

			assert.Equal(t, expectedBytes, actualBytes)
		})
	}
}
