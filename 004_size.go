package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(struct{}{}))
	var i interface{}
	fmt.Println(unsafe.Sizeof(i))
	var b bool
	fmt.Println(unsafe.Sizeof(b))
	// 64 arch:
	// 0
	// 0
	// 16
	// 1
	//
	// 32 arch
	// 0
	// 0
	// 8
	// 1
}
