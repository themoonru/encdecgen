package template

type MsgpackEncDecTemplateData struct {
	Package               string
	StructName            string
	StructNameFirstLetter string
	FieldsCount           int
	Fields                []Field
}

type Field struct {
	MsgPackType string
	Name        string
	Index       int64
}
