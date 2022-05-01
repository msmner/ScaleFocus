package main

import (
	api "Lecture27/Task/API"
	db "Lecture27/Task/db"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func main() {
	dsn := "root:Manotan88@/go_hackerrank?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	queries := db.New(dbConn)
	router := http.NewServeMux()
	router.HandleFunc("/top", api.StoriesHandler(queries))
	log.Fatal(http.ListenAndServe(":9000", router))
}
