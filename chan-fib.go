package main

import "fmt"

func main() {
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
