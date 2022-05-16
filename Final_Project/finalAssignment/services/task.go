package services

import (
	"final/models"
	"final/persistence"
	"log"
)

type TaskService struct {
	taskRepository *persistence.TaskRepository
}

func NewTaskService(tr *persistence.TaskRepository) *TaskService {
	return &TaskService{taskRepository: tr}
}

func (ts *TaskService) CreateTask(text string, listId int64, completed bool) (models.Task, error) {
	log.Printf("task service input parameters %s %d %v", text, listId, completed)
	task := models.Task{Text: text, ListID: int(listId), Completed: completed}
	id, err := ts.taskRepository.InsertTask(task)
	if err != nil {
		return task, err
	}

	log.Printf("id of inserted task %d", id)
	createdTask, err := ts.taskRepository.GetTask(id)
	if err != nil {
		return createdTask, err
	}

	return createdTask, nil
}

func (ts *TaskService) GetTasks(listId int64) ([]models.Task, error) {
	tasks, err := ts.taskRepository.GetTasks(listId)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (ts *TaskService) UpdateTask(taskId int64) (models.Task, error) {
	task, err := ts.taskRepository.UpdateTask(taskId)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (ts *TaskService) DeleteTask(id int64) error {
	err := ts.taskRepository.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}
