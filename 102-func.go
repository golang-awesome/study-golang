package main

import (
	"errors"
	"fmt"
)

type Printer func(content string) (n int, err error)

func printToStd(content string) (n int, err error) {
	return fmt.Println(content)
}

type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}

	return op(x, y), nil
}

type calculateFunc func(x, y int) (int, error)

func genCalculator(op func(x, y int) int) func(x, y int) (int, error) {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}

		return op(x, y), nil
	}
}

func identity(x int) int {
	return x
}

func main() {
	var p Printer
	fmt.Println(p)
	fmt.Println(p == nil)
	p = printToStd
	fmt.Println(p)
	p("something meow")
}
