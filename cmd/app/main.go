package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"get-post/internal/database"
	"get-post/internal/handlers"
	"get-post/internal/taskService"
	"get-post/internal/web/tasks"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.TaskService{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with err: %v", err)
	}
}
