package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"vac/features/user"
	"vac/features/user/presentation/request"
	"vac/features/user/presentation/response"
	"vac/helper"
	"vac/middleware"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) RegisterUserHandler(e echo.Context) error {
	userData := request.UserRequest{}

	err := e.Bind(&userData)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "Invalid payload data", err)
	}
	err = uh.userService.RegisterUser(userData.ToUserCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (uh *UserHandler) GetUsersHandler(e echo.Context) error {
	claims := middleware.ExtractClaim(e)
	role := claims["role"].(string)
	if role != "admin" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}

	data, err := uh.userService.GetUsers(user.UserCore{})
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToUserResponseList(data))
}

func (uh *UserHandler) LoginUserHandler(e echo.Context) error {
	userAuth := request.UserAuth{}
	e.Bind(&userAuth)

	data, err := uh.userService.LoginUser(userAuth.ToUserCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.TouserLoginResponse(data))
}

func (uh *UserHandler) GetUserByIdHandler(e echo.Context) error {
	claims := middleware.ExtractClaim(e)
	role := claims["role"].(string)
	if role != "admin" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}

	data, err := uh.userService.GetUserById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToUserResponse(data))
}
func (uh *UserHandler) GetUserOwnIdentity(e echo.Context) error {
	claims := middleware.ExtractClaim(e)
	role := claims["role"].(string)
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}
	userId := int(claims["id"].(float64))

	data, err := uh.userService.GetUserById(userId)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToUserResponse(data))
}

func (uh *UserHandler) UpdateUserHandler(e echo.Context) error {
	var userData request.UserRequest
	err := e.Bind(&userData)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}
	claims := middleware.ExtractClaim(e)
	userId := int(claims["id"].(float64))
	role := claims["role"]
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusBadRequest, "role not allowed to do action", errors.New("not allowed"))
	}
	userData.ID = userId

	err = uh.userService.UpdateUser(userData.ToUserCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}
