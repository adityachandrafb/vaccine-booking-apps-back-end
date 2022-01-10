package presentation

import (
	"net/http"
	"strconv"
	"vac/features/admin"
	"vac/features/admin/presentation/request"
	"vac/features/admin/presentation/response"
	"vac/helper"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService admin.Service
}

func NewAdminHandler(data admin.Service)*AdminHandler{
	return &AdminHandler{data}
}

func (ah *AdminHandler)RegisterAdminHandler(e echo.Context)error{
	reqData:=request.AdminRequest{}

	err:=e.Bind(&reqData)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}

	err=ah.AdminService.RegisterAdmin(request.FromAdminRequest(reqData))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ah *AdminHandler)LoginAdminHandler(e echo.Context)error{
	var adminLogin request.AdminLogin

	err:=e.Bind(&adminLogin)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}
	data, err:=ah.AdminService.LoginAdmin(request.FromAdminLogin(adminLogin))
	if err!=nil{
		return helper.ErrorResponse(e,http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToAdminLoginResponse(data))
}

func (ah *AdminHandler)GetAdminsHandler(e echo.Context)error{
	data,err:=ah.AdminService.GetAdmins()
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToAdminResponseList(data))
}

func (ah *AdminHandler)GetAdminByIdHandler(e echo.Context)error{
	id, err:=strconv.Atoi(e.Param("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	data, err:=ah.AdminService.GetAdminById(admin.AdminCore{ID:uint(id)})

	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToAdminResponse(data))

}