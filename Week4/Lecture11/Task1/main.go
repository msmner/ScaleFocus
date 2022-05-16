package main

import (
	"fmt"
	"sync"
)

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}

	evenCh := processEven(inputs)
	for evenVal := range evenCh {
		fmt.Println("Even Value: ", evenVal)
	}

	oddCh := processOdd(inputs)
	for oddVal := range oddCh {
		fmt.Println("Odd Value : ", oddVal)
	}
}

func processEven(inputs []int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, v := range inputs {
			wg.Add(1)
			go func(input int) {
				defer wg.Done()
				if input%2 == 0 {
					out <- input
				}
			}(v)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func processOdd(inputs []int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, v := range inputs {
			wg.Add(1)
			go func(input int) {
				defer wg.Done()
				if input%2 != 0 {
					out <- input
				}
			}(v)
		}
		wg.Wait()
		close(out)
	}()

	return out
}
