package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func New(name string) Dog {
	return Dog{name}
}

func TestSlicePointer(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr)

	sli := arr[:]
	fmt.Printf("%p\n", sli)
	fmt.Printf("%p\n", &sli[0])

	fmt.Printf("%p\n", &sli)          // addr of slice
	fmt.Println(unsafe.Pointer(&sli)) // addr of slice

	sliHeader := (*reflect.SliceHeader)(unsafe.Pointer(&sli))
	fmt.Printf("0x%10x\n", sliHeader.Data)
}

func main() {
	dog := New("little dog")
	dogP := &dog
	dog.SetName("monster")
	fmt.Println(dog)
	dogPtr := uintptr(unsafe.Pointer(dogP))
	fmt.Println(dogP, dogPtr)
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Println(nameP, *nameP)
	fmt.Println()
	fmt.Println()
	TestSlicePointer(nil)
}
