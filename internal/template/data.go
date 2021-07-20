package template

type MsgpackEncDecTemplateData struct {
	Package               string
	StructName            string
	StructNameFirstLetter string
	FieldsCount           int
	Fields                []Field
}

type Field struct {
	ConvType    string
	MsgPackType string
	Name        string
	Index       int64
}
