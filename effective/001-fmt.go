package main

import "fmt"

// go fmt
// gofmt

func main() {
	var x, y int
	z := x<<8 + y<<16
	fmt.Println(z)

	for i, v := range [...]int{1, 2, 3} {
		fmt.Println(i, "->", v)
	}
}
