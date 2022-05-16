package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type DbClient struct{}

func NewDbClient() *DbClient {
	return &DbClient{}
}

func (db *DbClient) Connect() (*sql.DB, error) {
	host := os.Getenv("host")
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		return nil, fmt.Errorf("error getting port %w", err)
	}
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbConn, err := sql.Open(user, psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database %w", err)
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database %w", err)
	}

	fmt.Println("Successfully connected!")

	return dbConn, nil
}
