//  go test -benchtime=10s -bench=. ./020-pool_test.go
package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"
)

func main() {
}

func init() {
	daemonStarted := startNetworkDeamonWithCache()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v\n", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v\n", err)
		}

		conn.Close()
	}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}

	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

func startNetworkDeamonWithCache() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}

		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v\n", conn)
				continue
			}

			svcConn := connPool.Get()
			fmt.Fprintf(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func connectToService() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()
			fmt.Fprintf(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func DemoCreate2() {
	var numCalsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// seed the pool with 4kb
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// quick done with the memory
			//time.Sleep(time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created", numCalsCreated)
}

func DemoCreate() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}
