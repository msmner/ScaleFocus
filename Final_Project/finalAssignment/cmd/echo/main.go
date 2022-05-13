package main

import (
	"database/sql"
	"final/cmd"
	"final/controllers"
	"final/persistence"
	"final/services"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "todo"
)

func main() {
	router := echo.New()
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	listRepository := persistence.NewListRepository(db)
	taskRepository := persistence.NewTaskRepository(db)
	listService := services.NewListService(listRepository)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)
	listController := controllers.NewListController(listService)

	// Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})

	//list routes
	router.GET("/api/lists", listController.GetLists)
	router.POST("/api/lists", listController.CreateList)
	router.DELETE("/api/lists/:id", listController.DeleteList)

	//task routes
	router.POST("/api/lists/:id/tasks", taskController.CreateTask)
	router.GET("/api/lists/:id/tasks", taskController.GetTasks)
	router.PATCH("/api/tasks/:id", taskController.UpdateTask)
	router.DELETE("/api/tasks/:id", taskController.DeleteTask)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
