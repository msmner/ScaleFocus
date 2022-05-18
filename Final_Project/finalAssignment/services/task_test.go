package services

import (
	"final/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockTaskRepository := MockTaskRepository{}
	taskService := NewTaskService(&mockTaskRepository)

	_, err := taskService.CreateTask("insert", 1, false)
	assert.EqualError(t, err, "error inserting task: error")

	_, err = taskService.CreateTask("error2", 2, false)
	assert.EqualError(t, err, "error getting task: error")

	expectedTask := models.Task{ID: 2, Text: "ok", ListID: 1, Completed: false}
	task, _ := taskService.CreateTask("ok", 1, false)
	assert.Equal(t, expectedTask, task)
}

func TestGetTasks(t *testing.T) {
	mockTaskRepository := MockTaskRepository{}
	taskService := NewTaskService(&mockTaskRepository)

	_, err := taskService.GetTasks(1)
	assert.EqualError(t, err, "error getting tasks: error")

	expectedTasks := []models.Task{{ID: 1, Text: "test", ListID: 2, Completed: false}, {ID: 2, Text: "test2", ListID: 2, Completed: false}}
	tasks, _ := taskService.GetTasks(2)
	assert.Equal(t, expectedTasks, tasks)
}

func TestUpdateTask(t *testing.T) {
	mockTaskRepository := MockTaskRepository{}
	taskService := NewTaskService(&mockTaskRepository)

	_, err := taskService.UpdateTask(1)
	assert.EqualError(t, err, "error updating task: error")

	expectedTask := models.Task{ID: 2, Text: "test", ListID: 2, Completed: false}
	task, _ := taskService.UpdateTask(2)
	assert.Equal(t, expectedTask, task)
}

func TestDeleteTask(t *testing.T) {
	mockTaskRepository := MockTaskRepository{}
	taskService := NewTaskService(&mockTaskRepository)

	err := taskService.DeleteTask(1)
	assert.EqualError(t, err, "error deleting task: error")
}
