package main

import (
	"fmt"
	"sync"
)

func main() {
	//TestChanType()
	//TestChanBlocking()

	//TestChanRange()
	//TestChanCloseCheap()

	// deadlock
	//TestChanDeadlock()

	//TestBufferedChan()
}

func TestBufferedChan() {
	//var ch1 chan interface{}
	//ch1 = make(chan interface{}, 4)
	ch1 := make(chan rune, 4)
	ch1 <- 'A'
	ch1 <- 'B'
	ch1 <- 'C'
	ch1 <- 'D'
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

func TestChanCloseCheap() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines")
	close(begin)
	wg.Wait()
}

func TestChanDeadlock() {
	strChan := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		strChan <- "hello channels"
	}()
	fmt.Println(<-strChan)
}

func TestChanRange() {
	ch := make(chan int)
	go func() {
		// ensure channel is closed before exiting the goroutine
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("ch val: %v\n", v)
	}
}

func TestChanBlocking() {
	intChan := make(chan int)
	go func() {
		intChan <- 12
	}()
	i, ok := <-intChan
	fmt.Println(i, ok)
	close(intChan)
	i, ok = <-intChan
	fmt.Println(i, ok)
	i, ok = <-intChan
	fmt.Println(i, ok)
}

func TestChanType() {
	// bidirectional
	var ch chan interface{}
	ch = make(chan interface{}, 12)
	// can only read
	var receiveChan <-chan interface{}
	receiveChan = make(<-chan interface{})
	fmt.Println(receiveChan)
	// can only send
	var sendChan chan<- interface{}
	sendChan = make(chan<- interface{})
	fmt.Println(sendChan)
	// valid
	sendChan = ch
	receiveChan = ch
}
