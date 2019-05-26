package main

import (
	"fmt"
	"sync"
)

func main() {
	TestDeadlock()
}

func TestDeadlock() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}

func TestOnce2() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }
	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
	// Count: 1
}

func TestOnce() {
	var count int
	increment := func() {
		count++
	}
	var once sync.Once

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}

	wg.Wait()
	fmt.Printf("Count is %d\n", count)
	// Count is 1
}
