package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	m int8
	n int64
}

func main() {
	var n int
	fmt.Println(unsafe.Sizeof(n))
	var t T
	fmt.Println(unsafe.Sizeof(t))
}
