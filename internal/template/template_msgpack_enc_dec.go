package template

const MsgpackEncDecTmpl = `
func ({{ .StructNameFirstLetter }} *{{ .StructName }}) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen({{ .FieldsCount }}); err != nil {
		return err
	}
	{{range .Fields }}
	if err := e.Encode{{ .MsgPackType }}({{ $.StructNameFirstLetter }}.{{ .Name }}); err != nil { // {{ .Index }}
		return err
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
