package main

import (
	"fmt"
	"unsafe"
)

/*
The empty struct and empty interface, though syntactically similar, are actually opposites. An empty struct holds no data; an empty interface can hold any type of value. If I see a map[MyType]struct{}, I know immediately that no values will be stored, only keys. If I see a map[MyType]interface{}, my first impression will be that it is a heterogenous collection of values. Even if I see code storing nil in it, I won't know for sure that some other piece of code doesn't store something else in it.
*/
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
