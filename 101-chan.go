package main

import (
	"fmt"
	"math/rand"
)

type Notifier interface {
	SendInt(ch chan<- int)
}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}

	close(ch)
	return ch
}

func testSelect() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("the index: %d\n", index)
	intChannels[index] <- index

	select {
	case <-intChannels[0]:
		fmt.Println("first")
	case <-intChannels[1]:
		fmt.Println("second")
	case elem := <-intChannels[2]:
		fmt.Println("third:", elem)
	default:
		fmt.Println("default")
	}
}

func main() {
	ch1 := make(chan int, 3)
	SendInt(ch1)
	SendInt(ch1)
	SendInt(ch1)
	elem1, close := <-ch1
	fmt.Println(elem1, close)
	elem2, close := <-ch1
	fmt.Println(elem2, close)
	elem3, close := <-ch1
	fmt.Println(elem3, close)

	ch := make(chan []int, 1)
	s1 := []int{1, 2, 3}
	ch <- s1
	s2 := <-ch
	s2[0] = 100
	fmt.Println(s1, s2)

	ch2 := make(chan [2]int, 1)
	s3 := [...]int{1, 2}
	ch2 <- s3
	s4 := <-ch2
	s3[0] = 100
	fmt.Println(s3, s4)

	var uselessChan = make(chan<- []int, 4)
	fmt.Println(uselessChan)
	var uselessChan2 = make(<-chan []int, 4)
	fmt.Println(uselessChan2)

	ch3 := getIntChan()

	for elem := range ch3 {
		fmt.Println(elem)
	}

	fmt.Println()
	fmt.Println()
	testSelect()
}
