package factory

import (
	"vac/driver"

	userData "vac/features/user/data"
	userPresent "vac/features/user/presentation"
	userService "vac/features/user/service"
)

type vacPresenter struct{
	UserPresentation userPresent.UserHandler
}

func Init() vacPresenter{
	userData:=userData.NewMysqlUserRepository(driver.DB)
	userService:=userService.NewUserService(userData)

	return vacPresenter{
		UserPresentation: *userPresent.NewUserHandler(userService),
	}
}