package api

import (
	db "Lecture27/Task/db"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	template "text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SimpleHandlerResponse struct {
}

type StoriesResponse struct {
	TopStories []Story `json:"top_stories"`
}

type Story struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Score      int       `json:"score"`
	Created_At time.Time `json:"created_at"`
}

type PageData struct {
	PageTitle string
	Links     []Story
}

func StoriesHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stories := make([]Story, 0, 10)
		ctx := context.Background()

		rows, err := queries.SelectRecords(ctx)
		if err != nil {
			fmt.Printf("Error querying the database: %v", err)
		}

		if len(rows) == 0 {
			stories, err = GenerateAndInsertRecords(stories, queries, ctx)
			if err != nil {
				fmt.Printf("Error inserting stories: %v", err)
			}
		} else {
			for _, v := range rows {
				lastHour := time.Now().Add(-1 * time.Second)
				if !lastHour.Before(v.CreatedAt) {
					err := queries.DeleteStory(ctx)
					if err != nil {
						fmt.Printf("Error deleting records: %v", err)
					}

					stories, err = GenerateAndInsertRecords(stories, queries, ctx)
					if err != nil {
						fmt.Printf("Error inserting stories: %v", err)
					}
					break
				}

				for _, story := range rows {
					s := Story{
						ID:         int(story.StoryID),
						Title:      story.StoryTitle,
						Score:      int(story.StoryScore),
						Created_At: story.CreatedAt,
					}
					stories = append(stories, s)
				}
			}
		}

		tmpl := template.Must(template.ParseFiles("top.html"))
		data := PageData{
			PageTitle: "Top Stories",
			Links:     stories,
		}
		tmpl.Execute(w, data)
	}
}

func GenerateAndInsertRecords(stories []Story, queries *db.Queries, ctx context.Context) ([]Story, error) {
	storiesIds := GetStoriesIds()
	storiesChannel := GenerateStoryDetails(storiesIds)
	for s := range storiesChannel {
		stories = append(stories, s)
		_, err := queries.CreateStory(ctx, db.CreateStoryParams{
			StoryID:    int32(s.ID),
			StoryTitle: s.Title,
			StoryScore: int32(s.Score),
			CreatedAt:  s.Created_At,
		})
		if err != nil {
			return nil, err
		}
	}

	return stories, nil
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
