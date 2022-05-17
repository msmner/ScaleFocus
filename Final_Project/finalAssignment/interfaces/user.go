package interfaces

import "final/models"

type IUserRepository interface {
	GetUser(username string) (models.User, error)
	AddListIdToUser(username string, listId int) error
	DeleteListFromUser(listId int, username string) error
}

type IUserService interface {
	GetUser(username string) (models.User, error)
}
