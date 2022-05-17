package interfaces

import "final/models"

type IListRepository interface {
	GetLists(user models.User) ([]models.List, error)
	DeleteList(id int) error
	InsertList(name string) (models.List, error)
}

type IListService interface {
	CreateList(name string, username interface{}) (models.List, error)
	GetLists(username interface{}) ([]models.List, error)
	DeleteList(username interface{}, listId string) error
}
