package interfaces

import "final/models"

type ITaskRepository interface {
	GetTasks(listId int) ([]models.Task, error)
	DeleteTask(id int) error
	UpdateTask(id int) (models.Task, error)
	GetTask(id int) (models.Task, error)
	InsertTask(task models.Task) (int, error)
}

type ITaskService interface {
	CreateTask(text string, listId int, completed bool) (models.Task, error)
	GetTasks(listId int) ([]models.Task, error)
	UpdateTask(taskId int) (models.Task, error)
	DeleteTask(id int) error
}
