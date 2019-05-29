package main

import (
	"container/list"
	"fmt"
	"runtime/debug"
)

type ApiError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) ApiError {
	return ApiError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err ApiError) Error() string {
	return err.Message
}

func main() {
	var ls list.List
	ls.PushBack("end")
	ls.PushBack("end")
	ls.PushBack("end")
	ls.PushBack("end")
	var x interface{}
	fmt.Println(x)
	fmt.Println(ls)
}
