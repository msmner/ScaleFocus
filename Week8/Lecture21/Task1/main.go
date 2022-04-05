package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/top", StoriesHandler())
	log.Fatal(http.ListenAndServe(":9000", router))
}

type SimpleHandlerResponse struct {
}

type StoriesResponse struct {
	TopStories []Story `json:"top_stories"`
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

func StoriesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stories := make([]Story, 0, 10)
		storiesIds := GetStoriesIds()
		storiesChannel := GenerateStoryDetails(storiesIds)
		for storyDetails := range storiesChannel {
			stories = append(stories, storyDetails)
			fmt.Println(storyDetails)
		}

		response := StoriesResponse{stories}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GenerateStoryDetails(data []string) <-chan Story {
	channel := make(chan Story, 10)
	defer close(channel)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json?print=pretty", data[index])
			body := ProcessRequest(url)
			var story Story
			err := json.Unmarshal(body, &story)
			if err != nil {
				fmt.Printf("Error deserializing the story: %v", err)
			}
			channel <- story
		}(i)
	}

	wg.Wait()
	return channel
}

func GetStoriesIds() []string {
	body := ProcessRequest("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	stringBody := strings.Trim(string(body), "[ ")
	result := strings.Split(stringBody, ", ")

	return result
}

func ProcessRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error building the request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error getting the response: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading the body: %v", err)
	}

	return body
}
