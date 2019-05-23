package _19_test

import (
	"runtime"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	t.Log("hello, i am a functional test...")
	t.Parallel()
}

func TestHello2(t *testing.T) {
	t.Parallel()
	t.Log("hello i am another functional test...")
}

func TestFail(t *testing.T) {
	t.Fail()
	t.Log("failed...")
	t.Log(runtime.GOMAXPROCS(0))
	t.Error("error")
	t.FailNow()
}

func testSimulate(max int) {
	time.Sleep(time.Millisecond * 50)
}

func BenchmarkSleep(b *testing.B) {
	b.StopTimer()
	time.Sleep(time.Millisecond * 5000)
	max := 1000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		testSimulate(max)
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log("hello i am benching marking")
	}
}
