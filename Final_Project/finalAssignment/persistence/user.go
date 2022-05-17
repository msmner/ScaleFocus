package persistence

import (
	"database/sql"
	"final/models"
	"fmt"
	"strconv"
	"strings"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(username string) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE username=$1`
	rows, err := r.db.Query(query, username)
	if err != nil {
		return user, fmt.Errorf("error in query getting user: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.PasswordHash, &user.ListIds)

		if err != nil {
			return user, fmt.Errorf("error in rows getting user: %w", err)
		}
	}

	err = rows.Err()
	if err != nil {
		return user, fmt.Errorf("error in rows getting user: %w", err)
	}
	return user, nil
}

func (ur *UserRepository) AddListIdToUser(username string, listId int) error {
	user, err := ur.GetUser(username)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}
	newlistId := fmt.Sprintf(",%d", listId)
	updatedListIds := user.ListIds + newlistId
	query := `UPDATE users SET listids=$1 WHERE username=$2`
	_, err = ur.db.Exec(query, updatedListIds, username)
	if err != nil {
		return fmt.Errorf("error adding list to user: %w", err)
	}

	return nil
}

func (ur *UserRepository) DeleteListFromUser(listId int, username string) error {
	user, err := ur.GetUser(username)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}
	listIds := strings.Trim(user.ListIds, ",")
	listIdsSlice := strings.Split(listIds, ",")
	for i, v := range listIdsSlice {
		id, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("error converting id string to id int delete list from user: %w", err)
		}
		if id == int(listId) {
			listIdsSlice[i] = listIdsSlice[len(listIdsSlice)-1]
			listIdsSlice[len(listIdsSlice)-1] = ""
			listIdsSlice = listIdsSlice[:len(listIdsSlice)-1]
			break
		}
	}

	newListIds := strings.Join(listIdsSlice, ",")
	query := `UPDATE users SET listIds=$1 WHERE username=$2`
	_, err = ur.db.Exec(query, newListIds, username)
	if err != nil {
		return fmt.Errorf("error updating user with new list id: %w", err)
	}
	return nil
}
