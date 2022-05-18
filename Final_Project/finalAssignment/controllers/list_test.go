package controllers

import (
	"final/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	list1JSON = `{"id":1,"name":"test1"}
`
	listsJSON = `[{"id":1,"name":"test1"},{"id":2,"name":"test2"}]
`
)

type MockListService struct{}

func (ls *MockListService) CreateList(name string, username interface{}) (models.List, error) {
	list := models.List{ID: 1, Name: "test1"}
	return list, nil
}

func (ls *MockListService) GetLists(username interface{}) ([]models.List, error) {
	list1 := models.List{ID: 1, Name: "test1"}
	list2 := models.List{ID: 2, Name: "test2"}
	listSlice := []models.List{list1, list2}
	return listSlice, nil
}

func (ls *MockListService) DeleteList(username interface{}, listId string) error {
	return nil
}

func TestCreateList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(list1JSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := MockListService{}
	listController := NewListController(&mockService)
	if assert.NoError(t, listController.CreateList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, list1JSON, rec.Body.String())
	}
}

func TestGetLists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := MockListService{}
	listController := NewListController(&mockService)

	if assert.NoError(t, listController.GetLists(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, listsJSON, rec.Body.String())
	}
}

func TestDeleteList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockService := MockListService{}
	listController := NewListController(&mockService)

	if assert.NoError(t, listController.DeleteList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"\"\n", rec.Body.String())
	}
}
