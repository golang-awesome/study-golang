package main

import "fmt"

func race() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

func main() {
	race()
	waitForever := make(chan interface{})
	go func() {
		panic("test panic")
	}()

	<-waitForever
}
