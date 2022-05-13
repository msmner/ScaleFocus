package services

import (
	"final/models"
	"final/persistence"
)

type ListService struct {
	listRepository *persistence.ListRepository
}

func NewListService(lr *persistence.ListRepository) *ListService {
	return &ListService{listRepository: lr}
}

func (ls *ListService) CreateList(name string) (models.List, error) {
	list := models.List{Name: name}
	id, err := ls.listRepository.InsertList(list)
	if err != nil {
		return list, err
	}

	createdList, err := ls.listRepository.GetList(id)
	if err != nil {
		return createdList, err
	}

	return createdList, nil
}

func (ls *ListService) GetLists() ([]models.List, error) {
	lists, err := ls.listRepository.GetLists()
	if err != nil {
		return lists, err
	}

	return lists, nil
}

func (ls *ListService) DeleteList(id int64) error {
	err := ls.listRepository.DeleteList(id)
	if err != nil {
		return err
	}

	return nil
}
