package main

import "fmt"

func addNum(num *int32, min int32, max int32, f func()) {
	defer f()
	if min > max {
		max = min
	}

	*num += max
	fmt.Printf("adding %d %d", min, max)
}

func coordinateWithChan() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)

	go addNum(&num, 1, max, func() {
		sign <- struct{}{}
	})

	go addNum(&num, 2, max, func() {
		sign <- struct{}{}
	})

	fmt.Println(len(sign))
	<-sign
	fmt.Println(len(sign))
	<-sign
	fmt.Println(len(sign))
}

func main() {
	coordinateWithChan()
}
