package main

import "fmt"

type User struct {
	Id       int    `msgpack_tnt_index:"3" json:"id"`
	Name     string `json:"name" msgpack_tnt_index:"1"`
	Password string `yaml:"password" msgpack_tnt_index:"2" json:"password"`
}

type Pswd struct {
	Value string `msgpack_tnt_index:"1" json:"value"`
	Salt  uint32 `json:"salt" msgpack_tnt_index:"2"`
}

const (
	CONST   = 1
	CONST_1 = "ненужный код"
)

func main() {
	a := 1
	b := "3"
	c := true
	f := bar(a, b, c)
	fmt.Println(f)
}

func bar(a int, b string, c bool) error {
	return nil
}
