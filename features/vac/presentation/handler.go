package presentation

import (
	"errors"
	"net/http"
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