package handlers

import (
	"context"
	"get-post/internal/userService"
	"get-post/internal/web/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error fetching users")
	}

	if len(allUsers) == 0 {
		return users.GetUsers200JSONResponse{}, nil
	}

	var response users.GetUsers200JSONResponse
	for _, usr := range allUsers {
		response = append(response, users.User{
			Id:    &usr.ID,
			Email: &usr.Email,
		})
	}

	return response, nil
}

func (u *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	if userRequest == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid user data")
	}

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
	}

	response := users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}
	return response, nil
}

func (u *UserHandler) DeleteUsersUserId(ctx context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	userId := request.UserId
	err := u.Service.DeleteUserByID(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	response := users.DeleteUsersUserId204Response{}
	return response, nil
}

func (u *UserHandler) PatchUsersUserId(ctx context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	userId := request.UserId
	userRequest := request.Body
	if userRequest == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid update data")
	}

	updatedUser := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUserInfo, err := u.Service.UpdateUserByID(userId, updatedUser)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	response := users.PatchUsersUserId200JSONResponse{
		Id:    &updatedUserInfo.ID,
		Email: &updatedUserInfo.Email,
	}
	return response, nil
}
