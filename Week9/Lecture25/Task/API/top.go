package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func StoriesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stories := make([]Story, 0, 10)

		dsn := "root:root@/go_hackerrank?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		query := `CREATE TABLE IF NOT EXISTS stories(story_id int not null primary key, story_title text not null,  
			story_score int not null, created_at datetime not null)`

		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		_, err = db.ExecContext(ctx, query)
		if err != nil {
			log.Printf("Error %s when creating product table", err)
		}

		rows, err := db.Query("select story_id, story_title, story_score, created_at from stories")
		if err != nil {
			fmt.Printf("Error querying the database: %v", err)
		}
		defer rows.Close()

		if !rows.Next() {
			stories = GenerateAndInsertRecords(stories, db)
		}

		for rows.Next() {
			var s Story
			rows.Scan(&s.ID, &s.Title, &s.Score, &s.Created_At)
			lastHour := time.Now().Add(-1 * time.Second)
			fmt.Printf("LAST HOUR IS :%v AND CREATED AT TIME IS: %v", lastHour, s.Created_At)
			if !lastHour.Before(s.Created_At) {
				deleteQuery := `DELETE FROM stories`
				res, _ := db.Exec(deleteQuery)
				rowsAff, _ := res.RowsAffected()
				fmt.Printf("ROWS AFFECTED DELETING RECORDS: %d", rowsAff)
				stories = GenerateAndInsertRecords(stories, db)
				break
			}

			stories = append(stories, s)
		}

		tmpl := template.Must(template.ParseFiles("top.html"))
		data := PageData{
			PageTitle: "Top Stories",
			Links:     stories,
		}
		tmpl.Execute(w, data)
	}
}

func GenerateAndInsertRecords(stories []Story, db *sql.DB) []Story {
	storiesIds := GetStoriesIds()
	storiesChannel := GenerateStoryDetails(storiesIds)
	for storyDetails := range storiesChannel {
		stories = append(stories, storyDetails)
		insertQuery := `INSERT INTO stories(story_id, story_title, story_score, created_at) VALUES (?, ?, ?, ?)`
		_, _ = db.Exec(insertQuery, storyDetails.ID, storyDetails.Title, storyDetails.Score, time.Now())
	}

	return stories
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
