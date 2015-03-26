package labs25

import (
	"runtime"
	"sync"
	"testing"
)

var c1 = make(chan int)
var c2 = make(chan int)

func init() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		for range c1 {
		}
	}()
	wg.Add(1)
	go func() {
		wg.Done()
		runtime.LockOSThread()
		for range c2 {
		}
	}()
	wg.Wait()
}

func Benchmark_Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c1 <- i
	}
}

func Benchmark_LockThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c2 <- i
	}
}
