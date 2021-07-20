package test

// map
//с ключами/значениями в виде скалярных типов (int, string, int32 и т.п.)
type MapScalarIntStruct struct {
	MapInt   map[int]int     `msgpack_index:"1"`
	MapInt8  map[int8]int8   `msgpack_index:"2"`
	MapInt16 map[int16]int16 `msgpack_index:"3"`
	MapInt32 map[int32]int32 `msgpack_index:"4"`
	MapInt64 map[int64]int64 `msgpack_index:"5"`
}

type MapScalarUIntStruct struct {
	MapUInt   map[uint]uint     `msgpack_index:"1"`
	MapUInt8  map[uint8]uint8   `msgpack_index:"2"`
	MapUInt16 map[uint16]uint16 `msgpack_index:"3"`
	MapUInt32 map[uint32]uint32 `msgpack_index:"4"`
	MapUInt64 map[uint64]uint64 `msgpack_index:"5"`
}

type MapScalarFloatStruct struct {
	MapFloat32 map[float32]float32 `msgpack_index:"1"`
	MapFloat64 map[float64]float64 `msgpack_index:"2"`
}

type MapScalarComplexStruct struct {
	MapComplex64  map[complex64]complex64   `msgpack_index:"1"`
	MapComplex128 map[complex128]complex128 `msgpack_index:"2"`
}

type MapScalarOtherStruct struct {
	MapByte   map[byte]byte     `msgpack_index:"1"`
	MapRune   map[rune]rune     `msgpack_index:"2"`
	MapString map[string]string `msgpack_index:"3"`
	MapBool   map[bool]bool     `msgpack_index:"4"`
}
