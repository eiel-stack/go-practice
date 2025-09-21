package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type S struct {
		f1 byte
		f2 struct{}
	}
	fmt.Println(unsafe.Sizeof(S{}))
}
