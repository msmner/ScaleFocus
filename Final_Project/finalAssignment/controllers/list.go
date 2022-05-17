package controllers

import (
	"final/models"
	"final/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ListController struct {
	listService *services.ListService
}

func NewListController(ls *services.ListService) *ListController {
	return &ListController{listService: ls}
}

func (lc *ListController) CreateList(c echo.Context) error {
	list := models.List{}
	if err := c.Bind(&list); err != nil {
		return err
	}

	user := c.Get("user")
	createdList, err := lc.listService.CreateList(list.Name, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdList)
}

func (lc *ListController) GetLists(c echo.Context) error {
	user := c.Get("user")
	lists, err := lc.listService.GetLists(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, lists)
}

func (lc *ListController) DeleteList(c echo.Context) error {
	id := c.Param("id")
	user := c.Get("user")
	err := lc.listService.DeleteList(user, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")
}
