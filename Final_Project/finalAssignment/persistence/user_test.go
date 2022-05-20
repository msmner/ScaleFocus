package persistence

import (
	"final/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	user := models.User{Username: "test", PasswordHash: "hash", ListIds: "1"}
	db, mock := NewMock()
	repo := &UserRepository{db}
	defer func() {
		repo.db.Close()
	}()

	query := "SELECT \\* FROM users WHERE username=\\$1"
	rows := sqlmock.NewRows([]string{"username", "passwordHash", "listIds"}).AddRow(user.Username, user.PasswordHash, user.ListIds)
	mock.ExpectQuery(query).WithArgs(user.Username).WillReturnRows(rows)
	user, err := repo.GetUser("test")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestAddListIdToUser(t *testing.T) {
	user := models.User{Username: "test", PasswordHash: "hash", ListIds: "1"}
	db, mock := NewMock()
	repo := &UserRepository{db}
	defer func() {
		repo.db.Close()
	}()

	getQuery := "SELECT \\* FROM users WHERE username=\\$1"
	rows := sqlmock.NewRows([]string{"username", "passwordHash", "listIds"}).AddRow(user.Username, user.PasswordHash, user.ListIds)
	mock.ExpectQuery(getQuery).WithArgs(user.Username).WillReturnRows(rows)
	updateQuery := "UPDATE users SET listids=\\$1 WHERE username=\\$2"
	mock.ExpectExec(updateQuery).WithArgs("1,1", "test").WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.AddListIdToUser("test", 1)

	assert.NoError(t, err)
}

func TestDeleteListFromUser(t *testing.T) {
	user := models.User{Username: "test", PasswordHash: "hash", ListIds: "1"}
	db, mock := NewMock()
	repo := &UserRepository{db}
	defer func() {
		repo.db.Close()
	}()

	getQuery := "SELECT \\* FROM users WHERE username=\\$1"
	rows := sqlmock.NewRows([]string{"username", "passwordHash", "listIds"}).AddRow(user.Username, user.PasswordHash, user.ListIds)
	mock.ExpectQuery(getQuery).WithArgs(user.Username).WillReturnRows(rows)
	updateQuery := "UPDATE users SET listIds=\\$1 WHERE username=\\$2"
	mock.ExpectExec(updateQuery).WithArgs("", "test").WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteListFromUser(1, "test")

	assert.NoError(t, err)
}
