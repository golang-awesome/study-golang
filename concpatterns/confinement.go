package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	TestDataPartition()
}

func TestDataPartition() {
	printData := func(wg *sync.WaitGroup, data [] byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}

func TestConfinement() {
	chanOwner := func() <-chan int {
		// new channel
		results := make(chan int, 5)

		// async write to channel
		// and then close the channel
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// read from channel
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}

func TestNonConfinement() {
	// data is available in both loopData & handleData channel
	data := []int{5, 2, 3, 4}
	loopData := func(handleData chan<- int) {
		defer close(handleData)

		for _, v := range data {
			handleData <- v
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
