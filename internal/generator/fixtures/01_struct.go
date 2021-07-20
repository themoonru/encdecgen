package test

type Struct01 struct {
	count int `msgpack_index:"1"`
}

func init() {
	_ = Struct01{count: 1}
}