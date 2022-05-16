package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var connections int
	var wg sync.WaitGroup

	flag.IntVar(&connections, "c", 2, "the number of connections to be processed concurrently")

	flag.Parse()
	if connections == 0 {
		flag.PrintDefaults()
		return
	}

	urls := flag.Args()
	bufferedChan := make(chan string, connections)

	for _, url := range urls {
		wg.Add(1)

		go func(connection string) {
			bufferedChan <- connection
			err := pingURL(connection)
			if err != nil {
				fmt.Printf("got error pinging url: %s", err)
			}
			<-bufferedChan
			wg.Done()
		}(url)
	}

	wg.Wait()
	close(bufferedChan)
}

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}
