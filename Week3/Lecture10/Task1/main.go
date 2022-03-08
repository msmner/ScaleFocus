package main

import (
	"fmt"
	"sync"
)

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.Print(times)
	cp.Wait()
}

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	state int
}

func (cp *ConcurrentPrinter) Print(times int) {
	for i := 0; i < 2*times; i++ {
		cp.WaitGroup.Add(1)
		go func() {
			defer cp.WaitGroup.Done()
			cp.Mutex.Lock()
			cp.state++
			if cp.state%2 == 0 {
				fmt.Print("bar")
			} else {
				fmt.Print("foo")
			}

			cp.Mutex.Unlock()
		}()
	}
}
