package persistence

import (
	"database/sql"
	"final/models"
	"log"
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
		err := rows.Scan(&user.Username, &user.PasswordHash, &user.ListId)
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

	log.Printf("user in updateuser persistence is %v", user)
	query := `INSERT INTO users (Username, PasswordHash, ListId) VALUES ($1, $2, $3)`
	_, err = ur.db.Exec(query, username, user.PasswordHash, listId)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteListFromUser(id int64, username string) error {
	query := `DELETE FROM users WHERE username=$1`
	_, err := ur.db.Exec(query, username)
	if err != nil {
		return err
	}
	return nil
}
