package handlers

import (
	"get-post/internal/userService"
	"get-post/internal/web/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if len(allUsers) == 0 {
		return ctx.JSON(http.StatusOK, []users.User{})
	}

	response := make([]users.User, len(allUsers))
	for i, u := range allUsers {
		id := u.ID
		email := u.Email
		password := u.Password

		response[i] = users.User{
			Id:       &id,
			Email:    &email,
			Password: &password,
		}
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) PostUsers(ctx echo.Context) error {
	var reqBody users.User
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON body",
		})
	}

	newUser := userService.User{
		Email:    *reqBody.Email,
		Password: *reqBody.Password,
	}

	createdUser, err := h.Service.CreateUser(newUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	id := createdUser.ID
	email := createdUser.Email
	password := createdUser.Password

	response := users.User{
		Id:       &id,
		Email:    &email,
		Password: &password,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (h *UserHandler) DeleteUsersUserId(ctx echo.Context, userId uint) error {
	err := h.Service.DeleteUserByID(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *UserHandler) PatchUsersUserId(ctx echo.Context, userId uint) error {
	var reqBody users.User
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON body",
		})
	}

	userUpdates := userService.User{
		Email:    *reqBody.Email,
		Password: *reqBody.Password,
	}

	updatedUser, err := h.Service.UpdateUserByID(userId, userUpdates)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	id := updatedUser.ID
	email := updatedUser.Email
	password := updatedUser.Password

	response := users.User{
		Id:       &id,
		Email:    &email,
		Password: &password,
	}

	return ctx.JSON(http.StatusOK, response)
}
