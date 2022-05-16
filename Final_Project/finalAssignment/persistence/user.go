package persistence

import (
	"database/sql"
	"final/models"
	"fmt"
	"log"
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
	log.Printf("username in getuser persistence is %s", username)
	user := models.User{}
	query := `SELECT * FROM users WHERE username=$1`
	rows, err := r.db.Query(query, username)
	if err != nil {
		log.Printf("rows err %v", err)
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.PasswordHash, &user.ListIds)

		if err != nil {
			return user, err
		}
	}

	err = rows.Err()
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) AddListIdToUser(username string, listId int64) error {
	user, err := ur.GetUser(username)
	if err != nil {
		return err
	}
	log.Printf("user in updateuser persistence is %v", username)
	newlistId := fmt.Sprintf(",%d", listId)
	updatedListIds := user.ListIds + newlistId
	log.Printf("updatedlist ids in persistence are %s", updatedListIds)
	query := `UPDATE users SET listids=$1 WHERE username=$2`
	_, err = ur.db.Exec(query, updatedListIds, username)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteListFromUser(listId int64, username string) error {
	user, err := ur.GetUser(username)
	if err != nil {
		return err
	}
	listIds := strings.Trim(user.ListIds, ",")
	listIdsSlice := strings.Split(listIds, ",")
	log.Printf("list ids slice in delete list from user persistence is %v", listIdsSlice)
	for i, v := range listIdsSlice {
		log.Printf("list id string from slice in delete list from user persistence is %v", v)
		id, err := strconv.Atoi(v)
		log.Printf("list id int from slice in delete list from user persistence is %v", id)

		if err != nil {
			return err
		}
		if id == int(listId) {
			listIdsSlice[i] = listIdsSlice[len(listIdsSlice)-1]
			listIdsSlice[len(listIdsSlice)-1] = ""
			listIdsSlice = listIdsSlice[:len(listIdsSlice)-1]
			break
		}
	}

	newListIds := strings.Join(listIdsSlice, ",")
	log.Printf("new list ids slice in delete list from user persistence is %v", newListIds)

	query := `UPDATE users SET listIds=$1 WHERE username=$2`
	_, err = ur.db.Exec(query, newListIds, username)
	if err != nil {
		return err
	}
	return nil
}
