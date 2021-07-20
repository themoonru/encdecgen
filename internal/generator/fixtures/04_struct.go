package test

type customInt int
type customInt8 int8
type customInt16 int16
type customInt32 int32
type customInt64 int64

type customUInt uint
type customUInt8 uint8
type customUInt16 uint16
type customUInt32 uint32
type customUInt64 uint64

type customFloat32 float32
type customFloat64 float64

type customComplex64 complex64
type customComplex128 complex128

type customByte byte
type customRune rune
type customString string
type customBool bool

//map
//с ключами/значением в виде кастомных типов на основе скалярных (type CustomString string)
type MapCustomIntStruct struct {
	MapInt   map[customInt]customInt     `msgpack_index:"1"`
	MapInt8  map[customInt8]customInt8   `msgpack_index:"2"`
	MapInt16 map[customInt16]customInt16 `msgpack_index:"3"`
	MapInt32 map[customInt32]customInt32 `msgpack_index:"4"`
	MapInt64 map[customInt64]customInt64 `msgpack_index:"5"`
}

type MapCustomUIntStruct struct {
	MapUInt   map[customUInt]customUInt     `msgpack_index:"1"`
	MapUInt8  map[customUInt8]customUInt8   `msgpack_index:"2"`
	MapUInt16 map[customUInt16]customUInt16 `msgpack_index:"3"`
	MapUInt32 map[customUInt32]customUInt32 `msgpack_index:"4"`
	MapUInt64 map[customUInt64]customUInt64 `msgpack_index:"5"`
}

type MapCustomFloatStruct struct {
	MapFloat32 map[customFloat32]customFloat32 `msgpack_index:"1"`
	MapFloat64 map[customFloat64]customFloat64 `msgpack_index:"2"`
}

type MapCustomComplexStruct struct {
	MapComplex64  map[customComplex64]customComplex64   `msgpack_index:"1"`
	MapComplex128 map[customComplex128]customComplex128 `msgpack_index:"2"`
}

type MapCustomOtherStruct struct {
	MapByte   map[customByte]customByte     `msgpack_index:"1"`
	MapRune   map[customRune]customRune     `msgpack_index:"2"`
	MapString map[customString]customString `msgpack_index:"3"`
	MapBool   map[customBool]customBool     `msgpack_index:"4"`
}
