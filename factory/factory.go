package factory

import (
	"vac/driver"

	userData "vac/features/user/data"
	userPresent "vac/features/user/presentation"
	userService "vac/features/user/service"

	adminData "vac/features/admin/data"
	adminPresent "vac/features/admin/presentation"
	adminService "vac/features/admin/service"
)

type vacPresenter struct{
	UserPresentation userPresent.UserHandler
	AdminPresentation adminPresent.AdminHandler

}

func Init() vacPresenter{
	userData:=userData.NewMysqlUserRepository(driver.DB)
	userService:=userService.NewUserService(userData)

	adminData:=adminData.NewAdminRepository(driver.DB)
	adminService:=adminService.NewAdminService(adminData)
	
	return vacPresenter{
		UserPresentation: *userPresent.NewUserHandler(userService),
		AdminPresentation: *adminPresent.NewAdminHandler(adminService),
	}
}