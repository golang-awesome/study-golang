package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//var nilChan chan interface{}

	// deadlock
	//<- nilChan

	// deadlock
	//nilChan <- struct{}{}

	// panic
	//close(nilChan)

	var stdoutBuff bytes.Buffer
	fmt.Println(stdoutBuff.Len(), stdoutBuff.Cap())
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Println(stdoutBuff.Len(), stdoutBuff.Cap())
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

func fibChan() {
	var fib func(n int) <-chan int
	fib = func(n int) <-chan int {
		// new channel
		result := make(chan int)

		// async write result
		go func() {
			defer close(result)
			if n <= 2 {
				result <- 1
				return
			}
			result <- <-fib(n - 1) + <-fib(n - 2)
		}()

		// return read only view
		return result
	}
	fmt.Println(<-fib(5))
}
