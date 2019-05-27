package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	testRandStream()
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

func testForRange() {
	stringStream := make(chan string, 3)
	done := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	for _, s := range []string{"a", "b", "c",} {
		select {
		case <-done:
			return
		case stringStream <- s:
			wg.Done()
		}
	}

	wg.Wait()
	close(stringStream)
	for v := range stringStream {
		fmt.Printf("%v\n", v)
	}
}

func testForSelectLoop() {
	for { // Either loop infinitely or range over something
		select {
		// do some work with channels
		}
	}
}

func testDoneSignal() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println("strings...")
					fmt.Println(s)
				case <-done:
					fmt.Println("done signaled...")
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		// cancel the op after 1 second
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

func testRandStream() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandomStream closure exited")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
					fmt.Println("generating one int")
				case <-done:
					fmt.Println("singal done")
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints")
	for i := 1; i < 4; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// simulate ongoing work
	time.Sleep(1 * time.Second)
	// 3 random ints
	// 1: 5577006791947779410
	// 2: 8674665223082153551
	// 3: 6129484611666145821
	// newRandomStream closure exited
}
