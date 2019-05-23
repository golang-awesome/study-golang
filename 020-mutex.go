package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func testSpinlock() {
	num2 := int32(10)
	for {
		if atomic.CompareAndSwapInt32(&num2, 10, 0) {
			fmt.Println("second number has gone to zero.")
			break
		}
	}

	time.Sleep(time.Millisecond * 500)
}

func main() {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("lock...")
	fmt.Println("unlocked...")
}
