package test

type CustomInt int
type CustomString string

type StructWithCustomTypes struct {
	customInt    CustomInt    `msgpack_index:"1"`
	customString CustomString `msgpack_index:"2"`
}

func init() {
	_ = StructWithCustomTypes{
		customInt:    1,
		customString: "1",
	}
}