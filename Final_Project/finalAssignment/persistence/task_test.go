package persistence

import (
	"final/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	task := models.Task{ID: 1, Text: "test", ListID: 1, Completed: false}
	db, mock := NewMock()
	repo := &TaskRepository{db}
	defer func() {
		repo.db.Close()
	}()

	query := "select \\* from tasks where listId=\\$1"
	rows := sqlmock.NewRows([]string{"id", "text", "listId", "completed"}).AddRow(task.ID, task.Text, task.ListID, task.Completed)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	tasks, err := repo.GetTasks(1)
	assert.NotNil(t, tasks)
	assert.NoError(t, err)
}

func TestDeleteTask(t *testing.T) {
	task := models.Task{ID: 1, Text: "test", ListID: 1, Completed: false}
	db, mock := NewMock()
	repo := &TaskRepository{db}
	defer func() {
		repo.db.Close()
	}()

	query := "DELETE FROM tasks WHERE ID=\\$1"
	mock.ExpectExec(query).WithArgs(task.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteTask(task.ID)

	assert.NoError(t, err)
}

func TestUpdateTask(t *testing.T) {
	task := models.Task{ID: 1, Text: "test", ListID: 1, Completed: false}
	db, mock := NewMock()
	repo := &TaskRepository{db}
	defer func() {
		repo.db.Close()
	}()

	getQuery := "SELECT \\* FROM tasks WHERE id=\\$1"
	rows := sqlmock.NewRows([]string{"id", "text", "listId", "completed"}).AddRow(task.ID, task.Text, task.ListID, task.Completed)
	mock.ExpectQuery(getQuery).WithArgs(task.ID).WillReturnRows(rows)
	updateQuery := "UPDATE tasks SET completed=\\$1 WHERE id=\\$2"
	mock.ExpectExec(updateQuery).WithArgs(true, task.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	task, err := repo.UpdateTask(task.ID)

	assert.NoError(t, err)
}

func TestGetTask(t *testing.T) {
	task := models.Task{ID: 1, Text: "test", ListID: 1, Completed: false}
	db, mock := NewMock()
	repo := &TaskRepository{db}
	defer func() {
		repo.db.Close()
	}()

	getQuery := "SELECT \\* FROM tasks WHERE id=\\$1"
	rows := sqlmock.NewRows([]string{"id", "text", "listId", "completed"}).AddRow(task.ID, task.Text, task.ListID, task.Completed)
	mock.ExpectQuery(getQuery).WithArgs(task.ID).WillReturnRows(rows)

	task, err := repo.GetTask(1)
	assert.NotNil(t, task)
	assert.NoError(t, err)
}

// func TestInsertTask(t *testing.T) {
// 	task := models.Task{ID: 1, Text: "test", ListID: 1, Completed: false}
// 	db, mock := NewMock()
// 	repo := &TaskRepository{db}
// 	defer func() {
// 		repo.db.Close()
// 	}()

// 	query := "INSERT INTO tasks\\(Text, ListId, Completed\\) VALUES\\(\\$1, \\$2, \\$3\\) RETURNING id"
// 	mock.ExpectExec(query).WithArgs(task.Text, task.ListID, task.Completed).WillReturnError(nil)
// 	id, err := repo.InsertTask(task)
// 	assert.NotNil(t, id)
// 	assert.NoError(t, err)
// }
