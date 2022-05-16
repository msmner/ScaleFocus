package persistence

import (
	"database/sql"
	"final/models"
	"fmt"
	"log"
	"strings"
)

type ListRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (r *ListRepository) GetLists(user models.User) ([]models.List, error) {
	log.Printf("user in getlists persistence is %v", user)
	lists := []models.List{}
	listIds := strings.Trim(user.ListIds, ",")
	listIdsSlice := strings.Split(listIds, ",")
	log.Printf("list ids after trimming are %s and after splitting are %v", listIds, listIdsSlice)
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
				return lists, err
			}
			lists = append(lists, list)
		}

		err = rows.Err()
		if err != nil {
			return lists, err
		}

	return lists, nil
}

func (r *ListRepository) DeleteList(id int64) error {
	deleteQuery := `DELETE FROM lists WHERE ID=$1`
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("error deleting %s", err)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *ListRepository) GetList(id int64) (models.List, error) {
	list := models.List{}
	query := `SELECT * FROM lists WHERE id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&list.ID, &list.Name)
		if err != nil {
			return list, err
		}
	}

	err = rows.Err()
	if err != nil {
		return list, err
	}
	return list, nil
}

func (r *ListRepository) InsertList(list models.List) (int64, error) {
	var id int64
	insertQuery := `INSERT INTO lists(Name) VALUES($1) RETURNING ID`
	err := r.db.QueryRow(insertQuery, list.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting new list: %w", err)
	}

	return id, nil
}
