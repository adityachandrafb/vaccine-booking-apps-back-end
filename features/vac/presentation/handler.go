package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"vac/features/vac"
	"vac/features/vac/presentation/request"
	"vac/features/vac/presentation/response"
	"vac/helper"
	"vac/middleware"

	"github.com/labstack/echo/v4"
)

type VacHandler struct {
	vacService vac.Service
}

func NewVacHandler(vs vac.Service) *VacHandler{
	return &VacHandler{vs}
}

func (vh *VacHandler) CreateVacPostHandler(e echo.Context) error{
	payloadData:=request.Vac{}
	err:=e.Bind(&payloadData)

	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payoad data", err)
	}
	claims:=middleware.ExtractClaim(e)
	role:=claims["role"].(string)
	if role!="admin"{
		return helper.ErrorResponse(e, http.StatusForbidden, "user not allowed to create job post", errors.New("not allowed"))
	}
	adminId:=claims["id"].(float64)
	payloadData.AdminId=int(adminId)
	err=vh.vacService.CreateVaccinationPost(payloadData.ToCore())

	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (vh *VacHandler) GetVacPostHandler(e echo.Context)error{
	var reqData request.VacFilter
	err:=e.Bind(reqData)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}
	data, err:= vh.vacService.GetVaccinationPost(reqData.ToCore())
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToVacResponseList(data))
}

func(vh *VacHandler)GetVacPostByIdHandler(e echo.Context)error{
	id, err:=strconv.Atoi(e.Param("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	data, err:=vh.vacService.GetVaccinationByIdPost(id)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToVacResponse(data))
}

func(vh *VacHandler)DeletVacPostHandler(e echo.Context)error{
	id, err:=strconv.Atoi(e.Param("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	claims:=middleware.ExtractClaim(e)
	adminId:=claims["id"].(float64)
	role:=claims["role"].(string)
	if role!="admin"{
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to delete data", errors.New("not allowed to delete data"))
	}
	err=vh.vacService.DeleteVaccinationPost(vac.VacCore{ID: id, AdminId: int(adminId)})
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}