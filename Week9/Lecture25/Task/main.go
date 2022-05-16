package main

import (
	api "Lecture25/Task/API"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/top", api.StoriesHandler())
	log.Fatal(http.ListenAndServe(":9000", router))
}
