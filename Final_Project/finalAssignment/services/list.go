package services

import (
	"final/interfaces"
	"final/models"
	"fmt"
	"strconv"
)

type ListService struct {
	listRepository interfaces.IListRepository
	userRepository interfaces.IUserRepository
}

func NewListService(lr interfaces.IListRepository, ur interfaces.IUserRepository) *ListService {
	return &ListService{listRepository: lr, userRepository: ur}
}

func (ls *ListService) CreateList(name string, username interface{}) (models.List, error) {
	list, err := ls.listRepository.InsertList(name)
	if err != nil {
		return list, fmt.Errorf("error creating list: %w", err)
	}

	err = ls.userRepository.AddListIdToUser(username.(string), list.ID)
	if err != nil {
		return list, fmt.Errorf("error adding list to user: %w", err)
	}

	return list, nil
}

func (ls *ListService) GetLists(username interface{}) ([]models.List, error) {
	user, err := ls.userRepository.GetUser(username.(string))
	if err != nil {
		return []models.List{}, fmt.Errorf("error getting user: %w", err)
	}
	lists, err := ls.listRepository.GetLists(user)
	if err != nil {
		return lists, fmt.Errorf("error getting lists: %w", err)
	}

	return lists, nil
}

func (ls *ListService) DeleteList(username interface{}, listId string) error {
	id, err := strconv.Atoi(listId)
	if err != nil {
		return fmt.Errorf("error converting listid: %w", err)
	}
	err = ls.listRepository.DeleteList(id)
	if err != nil {
		return fmt.Errorf("error deleting list: %w", err)
	}
	err = ls.userRepository.DeleteListFromUser(id, username.(string))
	if err != nil {
		return fmt.Errorf("error deleting list from user: %w", err)
	}

	return nil
}
