package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	testGOMAXPROCS()
}

func testGOMAXPROCS() {
	gomaxprocs := runtime.GOMAXPROCS(-1)
	fmt.Println(gomaxprocs)
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	previous := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(previous)
	// 8
	// 8
	// 8
}

func testDefault() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

func testTimeout() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Time out.")
		// default:
		// 	fmt.Println("default case")
		// default case
	}
}

func testCount() {
	c1 := make(chan interface{});
	close(c1)
	c2 := make(chan interface{});
	close(c2)

	// block forever
	// select {}

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
	// c1Count: 541
	// c2Count: 460
}

func testWait5Seconds() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(5 * time.Second)
	}()

	fmt.Println("Blocking on read...")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func testForLoop() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to Stop", workCounter)
}
