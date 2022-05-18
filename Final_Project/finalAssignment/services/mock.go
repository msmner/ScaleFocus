package services

import (
	"errors"
	"final/models"
)

type MockTaskRepository struct{}

func (ls *MockTaskRepository) GetTasks(listId int) ([]models.Task, error) {
	if listId == 1 {
		return nil, errors.New("error")
	}
	return []models.Task{{ID: 1, Text: "test", ListID: 2, Completed: false}, {ID: 2, Text: "test2", ListID: 2, Completed: false}}, nil
}

func (ls *MockTaskRepository) DeleteTask(id int) error {
	if id == 1 {
		return errors.New("error")
	}
	return nil
}

func (ls *MockTaskRepository) UpdateTask(id int) (models.Task, error) {
	if id == 1 {
		return models.Task{}, errors.New("error")
	}
	return models.Task{ID: 2, Text: "test", ListID: 2, Completed: false}, nil
}

func (ls *MockTaskRepository) GetTask(id int) (models.Task, error) {
	if id == 1 {
		return models.Task{}, errors.New("error")
	}
	return models.Task{ID: 2, Text: "ok", ListID: 1, Completed: false}, nil
}

func (ls *MockTaskRepository) InsertTask(task models.Task) (int, error) {
	if task.Text == "insert" {
		return 0, errors.New("error")
	}

	if task.Text == "ok" {
		return 2, nil
	}

	return 1, nil
}

type MockListRepository struct{}

func (ls *MockListRepository) GetLists(user models.User) ([]models.List, error) {
	if user.Username == "error2" {
		return nil, errors.New("error")
	}
	return []models.List{{ID: 1, Name: "test1"}, {ID: 2, Name: "test2"}}, nil
}

func (ls *MockListRepository) DeleteList(id int) error {
	if id == 1 {
		return errors.New("error")
	}
	return nil
}

func (ls *MockListRepository) InsertList(name string) (models.List, error) {
	if name == "error" {
		return models.List{}, errors.New("test error")
	}
	return models.List{ID: 1, Name: name}, nil
}

type MockUserRepository struct{}

func (ls *MockUserRepository) GetUser(username string) (models.User, error) {
	if username == "error" {
		return models.User{}, errors.New("error")
	}
	return models.User{Username: username, PasswordHash: "hash", ListIds: "16,17"}, nil
}

func (ls *MockUserRepository) AddListIdToUser(username string, listId int) error {
	if username == "testusername" {
		return errors.New("test error")
	}
	return nil
}

func (ls *MockUserRepository) DeleteListFromUser(listId int, username string) error {
	if listId == 2 {
		return errors.New("error")
	}
	return nil
}
