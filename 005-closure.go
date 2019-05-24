package main

import (
	"fmt"
	"sync"
)

func main() {
	testCreationCopy()
}

func testCreationCopy() {
	var wg sync.WaitGroup
	for _, hello := range []string{"hello", "world", "good day",} {
		wg.Add(1)
		go func(hello string) {
			defer wg.Done()
			fmt.Println(hello)
		}(hello)
	}
	wg.Wait()
}

func testCreation() {
	var wg sync.WaitGroup
	for _, hello := range []string{"hello", "world", "good day",} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(hello)
		}()
	}
	wg.Wait()
}

func testClosureCaptureReference() {
	var wg sync.WaitGroup
	hello := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		hello = "welcome"
	}()
	wg.Wait()
	fmt.Println(hello)
	// welcome
}
