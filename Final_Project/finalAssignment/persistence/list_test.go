package persistence

import (
	"database/sql"
	"final/models"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetLists(t *testing.T) {
	user := models.User{Username: "test", PasswordHash: "hash", ListIds: "1"}
	list := models.List{ID: 1, Name: "name"}
	db, mock := NewMock()
	repo := &ListRepository{db}
	defer func() {
		repo.db.Close()
	}()

	query := "select \\* from lists where id=\\$1"
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(list.ID, list.Name)
	mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows)
	lists, err := repo.GetLists(user)
	assert.NotNil(t, lists)
	assert.NoError(t, err)
}

func TestDeleteList(t *testing.T) {
	list := models.List{ID: 1, Name: "name"}
	db, mock := NewMock()
	repo := &ListRepository{db}
	defer func() {
		repo.db.Close()
	}()

	query := "DELETE FROM lists WHERE ID=\\$1"
	mock.ExpectExec(query).WithArgs(list.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteList(list.ID)

	assert.NoError(t, err)
}

// func TestInsertList(t *testing.T) {
// 	list := models.List{ID: 1, Name: "name"}
// 	db, mock := NewMock()
// 	repo := &ListRepository{db}
// 	defer func() {
// 		repo.db.Close()
// 	}()
// 	query := "INSERT INTO lists \\(Name\\) VALUES \\(\\$1\\) RETURNING ID"
// 	mock.ExpectExec(query).WithArgs(list.Name).WillReturnResult(sqlmock.NewResult(0, 1))
// 	list, err := repo.InsertList(list.Name)
// 	assert.NoError(t, err)
// }
