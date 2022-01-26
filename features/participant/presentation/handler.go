package presentation

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"vac/features/participant"
	"vac/features/participant/presentation/request"
	"vac/features/participant/presentation/response"
	"vac/helper"
	"vac/middleware"

	"github.com/labstack/echo/v4"
)

type ParticipantHandler struct {
	participantService participant.Service
}

func NewParticipantHandler(ps participant.Service) *ParticipantHandler {
	return &ParticipantHandler{ps}
}

func (ph *ParticipantHandler)UpdateParticipantHandler(e echo.Context) error{
	payloadData:=request.ParticipantUpdateRequest{}
	err:=e.Bind(&payloadData)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payload data", err)
	}
	claims:=middleware.ExtractClaim(e)
	payloadData.UserID=uint(claims["id"].(float64))
	role:=claims["role"].(string)
	if role!="user"{
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to update data", errors.New("not allowed"))
	}
	err=ph.participantService.UpdateParticipant(payloadData.ToCore())
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler)DeleteParticipantHandler(e echo.Context) error{
	id, err:=strconv.Atoi(e.Param("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest,"invalid id parameter", err)
	}
	claims:=middleware.ExtractClaim(e)
	userId:=claims["id"].(float64)
	role:=claims["role"].(string)
	if role!="user"{
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to delete data", errors.New("not allowed to delete data"))
	}
	err=ph.participantService.DeleteParticipant(participant.ParticipantCore{
		ID: uint(id),
		UserID: uint(userId),
	})
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler) ApplyParticipantHandler(e echo.Context) error {
	parData := request.ParticipantRequest{}
	err := e.Bind(&parData)

	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid payoad data", err)
	}
	vacId, err := strconv.Atoi(e.QueryParam("vacId"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "Invalid id parameter", err)
	}
	parData.VacID = uint(vacId)

	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	userId := uint(claims["id"].(float64))
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusForbidden, "you must login as user to apply participant", errors.New("not allowed"))
	}

	parData.UserID = userId
	err = ph.participantService.ApplyParticipant(parData.ToCore())

	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler) GetParticipantByUserIdHandler(e echo.Context) error {
	claims := middleware.ExtractClaim(e)
	userId := int(claims["id"].(float64))
	role := claims["role"].(string)
	if role != "user" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}

	participants, err := ph.participantService.GetParticipantByUserID(userId)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToParticipantResponseUserList(participants))
}

func (ph *ParticipantHandler) RejectParticipantHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	adminId := uint(claims["id"].(float64))
	if role != "admin" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to reject application", errors.New("action not allowed"))
	}
	err = ph.participantService.RejectParticipant(id, int(adminId))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler) AcceptParticipant(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	claims := middleware.ExtractClaim(e)
	role := claims["role"]
	adminId := uint(claims["id"].(float64))
	if role != "admin" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to accept participant", errors.New("action not allowed"))
	}
	err = ph.participantService.AcceptParticipant(id, int(adminId))

	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler) GetParticipantByIDHandler(e echo.Context) error {
	
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	parCore, err := ph.participantService.GetParticipantByID(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "somethinng went wrong", err)
	}
	
	claims := middleware.ExtractClaim(e)
	userId := uint(claims["id"].(float64))
	parCoreUserID :=parCore.UserID
	fmt.Println(userId)
	fmt.Println(parCore.User.ID)
	fmt.Println(parCore.UserID)

	if userId != parCoreUserID {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}
	return helper.SuccessResponse(e, response.ToParticipantResponse(parCore))
}

func (ph *ParticipantHandler) GetParticipantByVacIdHandler(e echo.Context) error {
	claims := middleware.ExtractClaim(e)
	role := claims["role"].(string)
	if role != "admin" {
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	par, err := ph.participantService.GetParticipantByVacID(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToParticipantResponseVacList(par))
}
