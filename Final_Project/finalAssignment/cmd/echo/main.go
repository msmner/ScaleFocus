package main

import (
	"final/auth"
	"final/cmd"
	"final/controllers"
	"final/persistence"
	"final/services"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading configuration variables %s", err)
	}

	dbClient := persistence.NewDbClient()
	db, err := dbClient.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database %s", err)
	}
	defer db.Close()

	router := echo.New()

	//basic authentication
	userRepository := persistence.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authMiddleware := auth.NewAuthMiddleware(userService)
	authMiddleware.Authenticate(router)

	//controllers and services
	listRepository := persistence.NewListRepository(db)
	taskRepository := persistence.NewTaskRepository(db)
	listService := services.NewListService(listRepository, userRepository)
	taskService := services.NewTaskService(taskRepository)
	taskController := controllers.NewTaskController(taskService)
	listController := controllers.NewListController(listService)
	weatherService := services.NewWeatherService()
	weatherController := controllers.NewWeatherController(weatherService)
	exportService := services.NewExportService(listRepository, taskRepository, userRepository)
	exportController := controllers.NewExportController(exportService)

	//list routes
	router.GET("/api/lists", listController.GetLists)
	router.POST("/api/lists", listController.CreateList)
	router.DELETE("/api/lists/:id", listController.DeleteList)
	router.GET("/api/list/export", exportController.ExportFile)

	//task routes
	router.POST("/api/lists/:id/tasks", taskController.CreateTask)
	router.GET("/api/lists/:id/tasks", taskController.GetTasks)
	router.PATCH("/api/tasks/:id", taskController.UpdateTask)
	router.DELETE("/api/tasks/:id", taskController.DeleteTask)

	//weather route
	router.GET("/api/weather", weatherController.GetWeather)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
