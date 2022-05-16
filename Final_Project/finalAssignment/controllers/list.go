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

	user := c.Get("user")
	log.Printf("create list controller user is %v", user)
	createdList, err := lc.listService.CreateList(list.Name, user)
	if err != nil {
		log.Printf("Error creating list: %v %v", err, createdList)
		return fmt.Errorf("error creating list: %w", err)
	}

	return c.JSON(http.StatusOK, createdList)
}

func (lc *ListController) GetLists(c echo.Context) (err error) {
	log.Printf("in get lists controller")
	user := c.Get("user")
	log.Printf("user in getlists controller is %v", user)
	lists, err := lc.listService.GetLists(user)
	if err != nil {
		log.Printf("Error getting lists: %v", err)
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
	log.Printf("Delete list id: %d", idInt)
	return c.JSON(http.StatusOK, "")
}
