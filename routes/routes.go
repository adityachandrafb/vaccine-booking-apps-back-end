package routes

import (
	"vac/config"
	"vac/factory"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo{
	presenter:=factory.Init()
	
	e := echo.New()
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	jwt.GET("/users", presenter.UserPresentation.GetUsersHandler)
	jwt.GET("/users/:id", presenter.UserPresentation.GetUserByIdHandler)
	e.POST("/users/register", presenter.UserPresentation.RegisterUserHandler)
	e.POST("/users/login", presenter.UserPresentation.LoginUserHandler)
	jwt.PUT("/users", presenter.UserPresentation.UpdateUserHandler)

	return e
}