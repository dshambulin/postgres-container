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

	// репозитории и сервисы
	tasksRepo := taskService.NewTaskRepository(database.DB)
	userRepo := userService.NewUserRepository(database.DB)

	tasksService := taskService.NewTaskService(tasksRepo)
	usrService := userService.NewUserService(userRepo)

	// хендлеры
	tasksHandler := handlers.NewTaskHandler(tasksService)
	userHandler := handlers.NewUserHandler(usrService)

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// рег. хендлеры обычным способом
	tasks.RegisterHandlers(e, tasksHandler)
	users.RegisterHandlers(e, userHandler)

	// запуск
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with err: %v", err)
	}
}
