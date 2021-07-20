package template

const MsgpackEncDecTmpl = `
func ({{ .StructNameFirstLetter }} *{{ .StructName }}) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen({{ .FieldsCount }}); err != nil {
		return fmt.Errorf("error encode struct {{ .StructName }} to MsgPack: %w", err)
	}
	{{range .Fields }}
	{{ if .ConvType }}if err := e.Encode{{ .MsgPackType }}({{ .ConvType }}({{ $.StructNameFirstLetter }}.{{ .Name }})); err != nil { // {{ .Index }}
		{{else}}if err := e.Encode{{ .MsgPackType }}({{ $.StructNameFirstLetter }}.{{ .Name }}); err != nil { // {{ .Index }}
		{{end}}return fmt.Errorf("error encode field {{ $.StructNameFirstLetter }}.{{ .Name }} to MsgPack: %w", err)
	}
	{{end}}
	return nil
}

func ({{ .StructNameFirstLetter }} *{{ .StructName }}) DecodeMsgpack(d *msgpack.Decoder) error {
	var err error
	var length int
	if length, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if length < {{ .FieldsCount }} {
		return fmt.Errorf("incorrect len: %d", length)
	}
	{{range .Fields }}
	if v, err := d.Decode{{ .MsgPackType }}(); err != nil { // {{ .Index }}
		return err
	} else {
		{{ $.StructNameFirstLetter }}.{{ .Name }} = v
	}
	{{end}}
	return nil
}
`
