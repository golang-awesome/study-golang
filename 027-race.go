package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%b\n%b\n", 124, 124 << 2, 124 * 4)
	var data int
	go func() {
		data++
	}()

	//time.Sleep(time.Millisecond)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}

	fmt.Println(12 &^ 11)
}
