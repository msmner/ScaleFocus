package services

import (
	"final/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateList(t *testing.T) {
	mockListRepository := MockListRepository{}
	mockUserRepository := MockUserRepository{}
	listService := NewListService(&mockListRepository, &mockUserRepository)

	list1, _ := listService.CreateList("test1", "testusername1")
	assert.Equal(t, models.List{ID: 1, Name: "test1"}, list1)

	_, err := listService.CreateList("error", "testusername")
	assert.EqualError(t, err, "error creating list: test error")

	_, err = listService.CreateList("test1", "testusername")
	assert.EqualError(t, err, "error adding list to user: test error")
}

func TestGetLists(t *testing.T) {
	mockListRepository := MockListRepository{}
	mockUserRepository := MockUserRepository{}
	listService := NewListService(&mockListRepository, &mockUserRepository)

	_, err := listService.GetLists("error")
	assert.EqualError(t, err, "error getting user: error")

	_, err = listService.GetLists("error2")
	assert.EqualError(t, err, "error getting lists: error")

	expectedLists := []models.List{{ID: 1, Name: "test1"}, {ID: 2, Name: "test2"}}
	lists, _ := listService.GetLists("ok")
	assert.Equal(t, expectedLists, lists)
}

func TestDeleteList(t *testing.T) {
	mockListRepository := MockListRepository{}
	mockUserRepository := MockUserRepository{}
	listService := NewListService(&mockListRepository, &mockUserRepository)

	err := listService.DeleteList("test", "a")
	assert.EqualError(t, err, "error converting listid: strconv.Atoi: parsing \"a\": invalid syntax")

	err = listService.DeleteList("test", "1")
	assert.EqualError(t, err, "error deleting list: error")

	err = listService.DeleteList("test", "2")
	assert.EqualError(t, err, "error deleting list from user: error")
}
