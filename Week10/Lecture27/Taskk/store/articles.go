package store

import (
	api "Lecture27/Task/API"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ArticlesRepository struct {
	db *sql.DB
}

func NewArticlesRepository(db *sql.DB) *ArticlesRepository {
	return &ArticlesRepository{db: db}
}

func (r *ArticlesRepository) CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS stories(story_id int not null primary key, story_title text not null,  
		story_score int not null, created_at datetime not null)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}

	return nil
}

func (r *ArticlesRepository) SelectRecords() (*sql.Rows, error) {
	rows, err := r.db.Query("select story_id, story_title, story_score, created_at from stories")
	if err != nil {
		fmt.Printf("Error querying the database: %v", err)
		return nil, err
	}

	return rows, nil
}

func (r *ArticlesRepository) DeleteRecords() error {
	deleteQuery := `DELETE FROM stories`
	_, err := r.db.Exec(deleteQuery)
	if err != nil {
		return err
	}

	return nil
}

func (r *ArticlesRepository) InsertRecords(story api.Story) error {
	insertQuery := `INSERT INTO stories(story_id, story_title, story_score, created_at) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(insertQuery, story.ID, story.Title, story.Score, time.Now())
	if err != nil {
		return err
	}

	return nil
}
