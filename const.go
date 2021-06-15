package main

import (
	"fmt"
	"os"
)

type Color int

const (
	Green Color = iota
	Red
	Blue
	Black
	Yellow
)

func foo() {
	fmt.Println("")
	os.Exit(0)
}

type Users struct {
	Name string `mbindex:"2"`
}
