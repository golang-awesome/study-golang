package main

import (
	"fmt"
	"sync"
)

func addNum(num *int32, min int32, max int32, f func()) {
	defer f()
	if min > max {
		max = min
	}

	*num += max
	fmt.Printf("adding %d %d\n", min, max)
}

//func addNum(num *int32, min int32, max int32, f func()) {
//	defer f()
//	if min > max {
//		max = min
//	}
//
//	*num += max
//
//}

func coordinateWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	fmt.Printf("The number: %d [with sync.WaitGroup]\n")
	max := int32(10)
	go addNum(&num, 3, max, func() {
		wg.Done()
	})
	go addNum(&num, 4, max, wg.Done)
	wg.Wait()
	fmt.Println("waited done")
	wg.Add(20)
}

func coordinateWithWaitGroupBatch() {
	total := 12
	stride := 4
	var num int32
	var wg sync.WaitGroup
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, int32(i+j), 12, wg.Done)
		}
		wg.Wait()
	}

	fmt.Println("End.")
}

func main() {
	coordinateWithWaitGroupBatch()
}
