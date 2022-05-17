package controllers

import (
	"final/interfaces"
	"final/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	taskService interfaces.ITaskService
}

func NewTaskController(ts interfaces.ITaskService) *TaskController {
	return &TaskController{taskService: ts}
}

func (tc *TaskController) CreateTask(c echo.Context) error {
	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}

	listIdStr := c.Param("id")
	listIdInt, err := strconv.Atoi(listIdStr)
	if err != nil {
		return err
	}

	createdTask, err := tc.taskService.CreateTask(task.Text, listIdInt, false)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdTask)
}

func (tc *TaskController) GetTasks(c echo.Context) error {
	listIdStr := c.Param("id")
	listIdInt, err := strconv.Atoi(listIdStr)
	if err != nil {
		return err
	}
	tasks, err := tc.taskService.GetTasks(listIdInt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c echo.Context) error {
	taskIdStr := c.Param("id")
	taskIdInt, err := strconv.Atoi(taskIdStr)
	if err != nil {
		return err
	}

	task, err := tc.taskService.UpdateTask(taskIdInt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c echo.Context) error {
	listIdStr := c.Param("id")
	listIdInt, err := strconv.Atoi(listIdStr)
	if err != nil {
		return err
	}

	tc.taskService.DeleteTask(listIdInt)

	return c.JSON(http.StatusOK, "")
}
