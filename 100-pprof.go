package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)
	log.SetPrefix("PPROF")

	go func() {
		goroutines := pprof.Lookup("goroutine")
		for range time.Tick(1 * time.Second) {
			log.Printf("goroutine count: %d\n", goroutines.Count())
		}
	}()

	var blockForever chan struct{}
	for i := 0; i < 10; i++ {
		go func() { <-blockForever }()
		time.Sleep(500 * time.Millisecond)
	}
}
