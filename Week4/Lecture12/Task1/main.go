package main

import (
	"log"
	"time"
)

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	bufferChan := make(chan string, bufferLimit)
	outChan := make(chan string)

	go func() {
		for {
			select {
			case bufferChan <- data:
				outChan <- data
			case <-time.After(clearInterval):
				for i := 0; i < bufferLimit; i++ {
					<-bufferChan
				}
			}
		}
	}()

	return outChan
}
