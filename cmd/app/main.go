package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"get-post/internal/database"
	"get-post/internal/handlers"
	"get-post/internal/taskService"
	"get-post/internal/userService"
	"get-post/internal/web/tasks"
	"get-post/internal/web/users"
)

func main() {
	database.InitDB()

	if database.DB == nil {
		log.Fatal("Database is not initialized")
	}

	if err := database.DB.AutoMigrate(&userService.User{}, &taskService.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	userRepo := userService.NewUserRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	userService := userService.NewUserService(userRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with err: %v", err)
	}
}
