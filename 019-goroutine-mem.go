package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)
	log.SetPrefix("PPROF")

	var c <-chan interface{}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutines := pprof.Lookup("goroutine")
		for range time.Tick(1 * time.Second) {
			log.Printf("goroutine count: %d\n", goroutines.Count())
		}
	}()

	noop := func() { wg.Done(); <-c }
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memConsumed()
	fmt.Println(after)
	fmt.Println(before)
	fmt.Println(numGoroutines)
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
