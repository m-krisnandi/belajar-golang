package main

import (
	"fmt"
	"sync"
	"runtime"
)

// deteksi race condition
// go run -race main.go

type counter struct {
	// cara 1 : menggunakan mutex 
	// sync.Mutex
	val int
}

func (c *counter) Add(int) {
	// cara 1 : menggunakan mutex
	// c.Lock()
	c.val++
	// c.Unlock()
}

func (c *counter) Value() int {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	// cara 2 : menggunakan mutex 
	var mtx sync.Mutex
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				// cara 2 : menggunakan mutex
				mtx.Lock()
				meter.Add(1)
				mtx.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.val)
	// atau
	// fmt.Println(meter.Value())
}