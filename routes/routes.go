package routes

import (
	"vac/config"
	"vac/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	e := echo.New()
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	jwt.GET("/users", presenter.UserPresentation.GetUsersHandler)
	jwt.GET("/user", presenter.UserPresentation.GetUserOwnIdentity)
	jwt.GET("/users/:id", presenter.UserPresentation.GetUserByIdHandler)
	e.POST("/users/register", presenter.UserPresentation.RegisterUserHandler)
	e.POST("/users/login", presenter.UserPresentation.LoginUserHandler)
	jwt.PUT("/users", presenter.UserPresentation.UpdateUserHandler)

	jwt.GET("/admins", presenter.AdminPresentation.GetAdminsHandler)
	jwt.GET("/admin/:id", presenter.AdminPresentation.GetAdminByIdHandler)
	e.POST("/admin/register", presenter.AdminPresentation.RegisterAdminHandler)
	e.POST("/admin/login", presenter.AdminPresentation.LoginAdminHandler)

	jwt.POST("/vac", presenter.VacPresentation.CreateVacPostHandler)
	e.GET("/vacs", presenter.VacPresentation.GetVacPostHandler)
	e.GET("/vac/:id", presenter.VacPresentation.GetVacPostByIdHandler)
	jwt.DELETE("/vac/:id", presenter.VacPresentation.DeletVacPostHandler)
	jwt.PUT("/vac", presenter.VacPresentation.UpdateVacPostHandler)
	e.POST("/near", presenter.VacPresentation.GetNearbyFacilitiesHandler)
	e.GET("/vacbyadmin/:id", presenter.VacPresentation.GetVacByIdAdminHandler)

	// participant
	jwt.POST("/participant", presenter.ParticipantPresentation.ApplyParticipantHandler)
	jwt.GET("/participant/:id", presenter.ParticipantPresentation.GetParticipantByIDHandler)
	jwt.GET("/participant/user", presenter.ParticipantPresentation.GetParticipantByUserIdHandler)
	jwt.GET("/participant/vac/:id", presenter.ParticipantPresentation.GetParticipantByVacIdHandler)
	jwt.PUT("/participant/canceled", presenter.ParticipantPresentation.RejectParticipantHandler)
	jwt.PUT("/participant/vaccinated", presenter.ParticipantPresentation.AcceptParticipant)

	jwt.PUT("/participant", presenter.ParticipantPresentation.UpdateParticipantHandler)

	jwt.DELETE("/participant/:id", presenter.ParticipantPresentation.DeleteParticipantHandler)
	return e
}
