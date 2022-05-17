package persistence

import (
	"database/sql"
	"final/models"
	"fmt"
	"strings"
)

type ListRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) GetLists(user models.User) ([]models.List, error) {
	lists := []models.List{}
	if user.ListIds != "" {
		listIds := strings.Trim(user.ListIds, ",")
		listIdsSlice := strings.Split(listIds, ",")
		for _, id := range listIdsSlice {
			rows, err := r.db.Query("select * from lists where id=$1", id)
			if err != nil {
				return lists, fmt.Errorf("error getting list from the database: %w", err)
			}

			defer rows.Close()
			for rows.Next() {
				list := models.List{}
				err := rows.Scan(&list.ID, &list.Name)
				if err != nil {
					return lists, fmt.Errorf("error scanning rows for get lists: %w", err)
				}
				lists = append(lists, list)
			}

			err = rows.Err()
			if err != nil {
				return lists, fmt.Errorf("rows error getting lists: %w", err)
			}
		}
	}

	return lists, nil
}

func (r *ListRepository) DeleteList(id int64) error {
	deleteQuery := `DELETE FROM lists WHERE ID=$1`
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("error executing query delete list: %w", err)
	}

	return nil
}

func (r *ListRepository) InsertList(name string) (models.List, error) {
	var id int
	insertQuery := `INSERT INTO lists(Name) VALUES($1) RETURNING ID`
	err := r.db.QueryRow(insertQuery, name).Scan(&id)
	list := models.List{ID: id, Name: name}
	if err != nil {
		return list, fmt.Errorf("error inserting new list: %w", err)
	}

	return list, nil
}
