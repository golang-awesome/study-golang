package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 10, 30, 1, 3, 4}
	go printCount(c)
}
