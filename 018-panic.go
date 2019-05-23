package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
	name string
}

func (error MyError) String() string {
	return fmt.Sprintf("%s: %s", error.name, error.error)
}

func (error MyError) Error() string {
	return error.String()
}

func main() {
	//arr := []int{1}
	// panic: runtime error: index out of range
	//fmt.Println(arr[2])

	fmt.Println("Enter func main")
	defer func() {
		fmt.Println("Enter defer func")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer func")
	}()

	panic(MyError{
		name:  "testnam",
		error: errors.New("test panic error"),
	})
	fmt.Println("Exit func main")
	//panic(12)
	//panic("wow a panic")
}
