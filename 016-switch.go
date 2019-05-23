package main

import "fmt"

func main() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch int8(1 + 3) {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}

	value6 := interface{}(byte(127))
	// byte is alias of uint8
	switch t := value6.(type) {
	case uint16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Println("Unsupported type: %T", t)
	}
}
