package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Date struct {
	i1 int
	i2 int
	i3 int
}

var global *Date
var mutex sync.Mutex

func Update() {
	d := new(Date)
	d.i1 = time.Now().Nanosecond()
	d.i2 = d.i1 + 1
	d.i3 = d.i1 + 2
	mutex.Lock()
	defer mutex.Unlock()
	global = d
}

func Get() *Date {
	mutex.Lock()
	defer mutex.Unlock()
	return global
}

func Process(c int, w *sync.WaitGroup) {
	d := Get()
	for i := 0; i < 10; i++ {
		fmt.Printf("c=%d i=%d i1=%d i2=%d i3=%d\n", c, i, d.i1, d.i2, d.i3)
		time.Sleep(time.Millisecond * time.Duration(rand.Int()%1000+500))
	}
	w.Done()
}

func main() {
	Update()
	var w sync.WaitGroup
	go func() {
		w.Add(1)
		for i := 0; i < 10000; i++ {
			Update()
			time.Sleep(time.Millisecond * 20)
		}
		w.Done()
	}()
	
	for i := 0; i < 100; i++ {
		go Process(i, &w)
		w.Add(1)
		time.Sleep(time.Millisecond*100)
	}



	w.Wait()
}
