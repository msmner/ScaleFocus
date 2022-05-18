package controllers

import (
	"final/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	task1JSON = `{"id":1,"text":"test1","listId":1,"completed":false}
`
	tasksJSON = `[{"id":1,"text":"test1","listId":1,"completed":false},{"id":2,"text":"test2","listId":1,"completed":false}]
`
)

type MockTaskService struct{}

func (ts *MockTaskService) CreateTask(text string, listId int, completed bool) (models.Task, error) {
	task := models.Task{ID: 1, Text: "test1", ListID: 1, Completed: false}
	return task, nil
}

func (ts *MockTaskService) GetTasks(listId int) ([]models.Task, error) {
	task1 := models.Task{ID: 1, Text: "test1", ListID: 1, Completed: false}
	task2 := models.Task{ID: 2, Text: "test2", ListID: 1, Completed: false}
	taskSlice := []models.Task{task1, task2}
	return taskSlice, nil
}

func (ts *MockTaskService) UpdateTask(taskId int) (models.Task, error) {
	task1 := models.Task{ID: 1, Text: "test1", ListID: 1, Completed: false}
	return task1, nil
}

func (ts *MockTaskService) DeleteTask(id int) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(task1JSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	mockService := MockTaskService{}
	taskController := NewTaskController(&mockService)
	if assert.NoError(t, taskController.CreateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, task1JSON, rec.Body.String())
	}
}

func TestGetTasks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	mockService := MockTaskService{}
	taskController := NewTaskController(&mockService)

	if assert.NoError(t, taskController.GetTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, tasksJSON, rec.Body.String())
	}
}

func TestDeleteTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	mockService := MockTaskService{}
	taskController := NewTaskController(&mockService)

	if assert.NoError(t, taskController.DeleteTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"\"\n", rec.Body.String())
	}
}

func TestUpdateTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	mockService := MockTaskService{}
	taskController := NewTaskController(&mockService)

	if assert.NoError(t, taskController.UpdateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, task1JSON, rec.Body.String())
	}
}
