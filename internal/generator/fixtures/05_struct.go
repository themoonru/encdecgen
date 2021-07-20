package test

//срезы и указатели ([]struct)
type StructPointer struct {
	Pointer *MapScalarIntStruct `msgpack_index:"1"`
}

type SlicePointer struct {
	Slice []*MapScalarIntStruct `msgpack_index:"1"`
}

type Slice struct {
	Int    []int     `msgpack_index:"1"`
	String []string  `msgpack_index:"2"`
	Float  []float64 `msgpack_index:"3"`
}

type SliceIntStruct struct {
	Slice []MapScalarIntStruct `msgpack_index:"1"`
}

type SliceFloatStruct struct {
	Slice []MapScalarFloatStruct `msgpack_index:"1"`
}
