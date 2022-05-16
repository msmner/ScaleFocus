package services

import (
	"final/models"
	"final/persistence"
)

type ListService struct {
	listRepository *persistence.ListRepository
	userRepository *persistence.UserRepository
}

func NewListService(lr *persistence.ListRepository, ur *persistence.UserRepository) *ListService {
	return &ListService{listRepository: lr, userRepository: ur}
}

func (ls *ListService) CreateList(name string, username interface{}) (models.List, error) {
	list := models.List{Name: name}
	id, err := ls.listRepository.InsertList(list)
	if err != nil {
		return list, err
	}

	err = ls.userRepository.AddListIdToUser(username.(string), id)
	if err != nil {
		return list, err
	}

	createdList, err := ls.listRepository.GetList(id)
	if err != nil {
		return createdList, err
	}

	return createdList, nil
}

func (ls *ListService) GetLists(username interface{}) ([]models.List, error) {
	user, err := ls.userRepository.GetUser(username.(string))
	if err != nil {
		return []models.List{}, err
	}
	lists, err := ls.listRepository.GetLists(user)
	if err != nil {
		return lists, err
	}

	return lists, nil
}

func (ls *ListService) DeleteList(username interface{}, listId int64) error {
	err := ls.listRepository.DeleteList(listId)
	if err != nil {
		return err
	}
	err = ls.userRepository.DeleteListFromUser(listId, username.(string))
	if err != nil {
		return err
	}

	return nil
}
