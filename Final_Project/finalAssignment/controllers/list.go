package controllers

import (
	"final/models"
	"final/services"
	"fmt"
	"log"
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

	createdList, err := lc.listService.CreateList(list.Name)
	if err != nil {
		log.Printf("Error creating list: %v %v", err, createdList)
		return fmt.Errorf("error creating list: %w", err)
	}

	return c.JSON(http.StatusOK, createdList)
}

func (lc *ListController) GetLists(c echo.Context) (err error) {
	lists, err := lc.listService.GetLists()
	if err != nil {
		log.Printf("Error getting lists: %v", err)
		return fmt.Errorf("error getting lists: %w", err)
	}
	return c.JSON(http.StatusOK, lists)
}

func (lc *ListController) DeleteList(c echo.Context) (err error) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	lc.listService.DeleteList(int64(idInt))
	log.Printf("Delete list id: %d", idInt)
	return c.JSON(http.StatusOK, "")
}
