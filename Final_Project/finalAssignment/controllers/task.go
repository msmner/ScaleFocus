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

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(ts *services.TaskService) *TaskController {
	return &TaskController{taskService: ts}
}

func (tc *TaskController) CreateTask(c echo.Context) (err error) {
	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}

	listIdStr := c.Param("id")
	listIdInt, err := strconv.Atoi(listIdStr)
	if err != nil {
		return err
	}

	createdTask, err := tc.taskService.CreateTask(task.Text, int64(listIdInt), false)
	if err != nil {
		return fmt.Errorf("error creating task %w", err)
	}

	return c.JSON(http.StatusOK, createdTask)
}

func (tc *TaskController) GetTasks(c echo.Context) (err error) {
	log.Printf("in get tasks controller")
	listIdStr := c.Param("id")
	listIdInt, err := strconv.Atoi(listIdStr)
	if err != nil {
		return err
	}
	tasks, err := tc.taskService.GetTasks(int64(listIdInt))
	if err != nil {
		return fmt.Errorf("error getting tasks: %w", err)
	}

	return c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c echo.Context) (err error) {
	taskIdStr := c.Param("id")
	taskIdInt, err := strconv.Atoi(taskIdStr)
	if err != nil {
		return err
	}

	task, err := tc.taskService.UpdateTask(int64(taskIdInt))
	if err != nil {
		return fmt.Errorf("error updating task %w", err)
	}

	return c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c echo.Context) (err error) {
	listIdStr := c.Param("id")
	listIdInt, err := strconv.Atoi(listIdStr)
	if err != nil {
		return err
	}

	tc.taskService.DeleteTask(int64(listIdInt))

	return c.JSON(http.StatusOK, "")
}
