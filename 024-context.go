package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

func addNum(num *int32, min int32, max int32, f func()) {
	defer f()
	if min > max {
		max = min
	}

	*num += 1
	fmt.Printf("adding %d %d\n", min, max)
}

func coordinateWithContext() {
	total := 12
	var num int32
	fmt.Println("Start.")
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, int32(i), 12, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}

	<-ctx.Done()
	fmt.Println("End.")

}
