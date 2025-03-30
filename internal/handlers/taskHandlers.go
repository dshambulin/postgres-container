package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"get-post/internal/taskService"
	"get-post/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(ctx echo.Context) error {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := make([]tasks.Task, 0, len(allTasks))
	for _, t := range allTasks {
		id := t.ID
		isDone := t.IsDone
		taskName := t.Task
		userId := t.UserID

		response = append(response, tasks.Task{
			Id:     &id,
			IsDone: &isDone,
			Task:   &taskName,
			UserId: &userId,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *TaskHandler) PostTasks(ctx echo.Context) error {
	var reqBody tasks.Task
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON body",
		})
	}

	newTask := taskService.Task{
		Task:   *reqBody.Task,
		IsDone: *reqBody.IsDone,
		UserID: *reqBody.UserId,
	}

	createdTask, err := h.Service.CreateTask(newTask)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	id := createdTask.ID
	isDone := createdTask.IsDone
	taskName := createdTask.Task
	userId := createdTask.UserID

	response := tasks.Task{
		Id:     &id,
		IsDone: &isDone,
		Task:   &taskName,
		UserId: &userId,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) DeleteTasksTaskId(ctx echo.Context, taskId uint) error {
	if err := h.Service.DeleteTaskByID(taskId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *TaskHandler) PatchTasksTaskId(ctx echo.Context, taskId uint) error {
	var reqBody tasks.Task
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON body",
		})
	}

	updatedData := taskService.Task{
		Task:   *reqBody.Task,
		IsDone: *reqBody.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskId, updatedData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	id := updatedTask.ID
	isDone := updatedTask.IsDone
	taskName := updatedTask.Task
	userId := updatedTask.UserID

	response := tasks.Task{
		Id:     &id,
		IsDone: &isDone,
		Task:   &taskName,
		UserId: &userId,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetUsersUserIdTasks(ctx echo.Context, userId uint) error {
	tasksForUser, err := h.Service.GetTasksByUserID(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if len(tasksForUser) == 0 {
		return ctx.JSON(http.StatusOK, []tasks.Task{})
	}

	response := make([]tasks.Task, 0, len(tasksForUser))
	for _, t := range tasksForUser {
		id := t.ID
		taskName := t.Task
		isDone := t.IsDone
		uId := t.UserID

		response = append(response, tasks.Task{
			Id:     &id,
			Task:   &taskName,
			IsDone: &isDone,
			UserId: &uId,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}
