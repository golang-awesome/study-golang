package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type Run func()

func hello2() {
	fmt.Println("hello")
}

func main() {
	go hello2()

	// anonymous
	go func() {
		fmt.Println("hello world")
	}()

	// variable
	sayHello := func() {
		fmt.Println("hello")
	}
	go sayHello()

	num := runtime.GOMAXPROCS(2)
	fmt.Println(num)
	num2 := runtime.GOMAXPROCS(2)
	fmt.Println(num2)
	fmt.Println()

	var count uint32
	fmt.Println("count=", count)
	trigger := func(i uint32, fn Run) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}

	trigger(10, func() {})
}
