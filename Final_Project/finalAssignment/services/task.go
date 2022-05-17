package services

import (
	"final/interfaces"
	"final/models"
	"fmt"
)

type TaskService struct {
	taskRepository interfaces.ITaskRepository
}

func NewTaskService(tr interfaces.ITaskRepository) *TaskService {
	return &TaskService{taskRepository: tr}
}

func (ts *TaskService) CreateTask(text string, listId int, completed bool) (models.Task, error) {
	task := models.Task{Text: text, ListID: int(listId), Completed: completed}
	id, err := ts.taskRepository.InsertTask(task)
	if err != nil {
		return task, fmt.Errorf("error inserting task: %w", err)
	}

	createdTask, err := ts.taskRepository.GetTask(id)
	if err != nil {
		return createdTask, fmt.Errorf("error getting task: %w", err)
	}

	return createdTask, nil
}

func (ts *TaskService) GetTasks(listId int) ([]models.Task, error) {
	tasks, err := ts.taskRepository.GetTasks(listId)
	if err != nil {
		return tasks, fmt.Errorf("error getting task: %w", err)
	}

	return tasks, nil
}

func (ts *TaskService) UpdateTask(taskId int) (models.Task, error) {
	task, err := ts.taskRepository.UpdateTask(taskId)
	if err != nil {
		return task, fmt.Errorf("error updating task: %w", err)
	}

	return task, nil
}

func (ts *TaskService) DeleteTask(id int) error {
	err := ts.taskRepository.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}

	return nil
}
