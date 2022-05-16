package controllers

import (
	"final/models"
	"final/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ListController struct {
	listService *services.ListService
}

func NewListController(ls *services.ListService) *ListController {
	return &ListController{listService: ls}
}

func (lc *ListController) CreateList(c echo.Context) (err error) {
	list := models.List{}
	if err := c.Bind(&list); err != nil {
		return err
	}

	user := c.Get("user")
	createdList, err := lc.listService.CreateList(list.Name, user)
	if err != nil {
		return fmt.Errorf("error creating list: %w", err)
	}

	return c.JSON(http.StatusOK, createdList)
}

func (lc *ListController) GetLists(c echo.Context) (err error) {
	user := c.Get("user")
	lists, err := lc.listService.GetLists(user)
	if err != nil {
		return fmt.Errorf("error getting lists: %w", err)
	}

	return c.JSON(http.StatusOK, lists)
}

func (lc *ListController) DeleteList(c echo.Context) (err error) {
	idStr := c.Param("id")
	user := c.Get("user")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	lc.listService.DeleteList(user, int64(idInt))
	return c.JSON(http.StatusOK, "")
}
