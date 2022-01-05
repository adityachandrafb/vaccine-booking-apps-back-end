package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"vac/features/participant"
	"vac/features/participant/presentation/response"
	"vac/helper"
	"vac/middleware"

	"github.com/labstack/echo/v4"
)

type ParticipantHandler struct {
	participantService participant.Service
}

func NewParticipantHandler(ps participant.Service)*ParticipantHandler{
	return &ParticipantHandler{ps}
}

func (ph *ParticipantHandler) ApplyParticipantHandler(e echo.Context)error{
	vacId, err:=strconv.Atoi(e.QueryParam("vacId"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "Invalid id parameter", err)
	}

	claims:=middleware.ExtractClaim(e)
	role:=claims["role"]
	userId:=uint(claims["id"].(float64))
	if role!="user"{
		return helper.ErrorResponse(e, http.StatusForbidden, "you must login as user to apply participant", err)
	}

	err=ph.participantService.ApplyParticipant(participant.ParticipantCore{
		VacID: uint(vacId),
		UserID: userId,
	})
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler)GetParticipantByUserIdHandler(e echo.Context)error{
	claims:=middleware.ExtractClaim(e)
	userId:=int(claims["id"].(float64))
	role:=claims["role"].(string)
	if role!="user"{
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to get data", errors.New("forbidden"))
	}

	participants, err:=ph.participantService.GetParticipantByUserID(userId)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, response.ToParticipantResponseUserList(participants))
}

func(ph *ParticipantHandler)RejectParticipantHandler(e echo.Context)error{
	id, err:=strconv.Atoi(e.QueryParam("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	claims:=middleware.ExtractClaim(e)
	role:=claims["role"]
	adminId:=uint(claims["id"].(float64))
	if role!="admin"{
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to reject application", errors.New("action not allowed"))
	}
	err=ph.participantService.RejectParticipant(id, int(adminId))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func(ph *ParticipantHandler)AcceptParticipant(e echo.Context)error{
	id, err:=strconv.Atoi(e.QueryParam("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	claims:=middleware.ExtractClaim(e)
	role:=claims["role"]
	adminId:=uint(claims["id"].(float64))
	if role!="admin"{
		return helper.ErrorResponse(e, http.StatusForbidden, "role not allowed to accept participant", errors.New("action not allowed"))
	}
	err=ph.participantService.AcceptParticipant(id, int(adminId))

	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}
	return helper.SuccessResponse(e, nil)
}

func (ph *ParticipantHandler)GetParticipantByIDHandler( e echo.Context)error{
	id, err:=strconv.Atoi(e.Param("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	parCore,err:=ph.participantService.GetParticipantByID(id)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "somethinng went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToParticipantResponse(parCore))
}

func (ph *ParticipantHandler)GetParticipantByVacIdHandler(e echo.Context)error{
	id, err:=strconv.Atoi(e.Param("id"))
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusBadRequest, "invalid id parameter", err)
	}
	par, err:=ph.participantService.GetParticipantByVacID(id)
	if err!=nil{
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", err)
	}

	return helper.SuccessResponse(e, response.ToParticipantResponseVacList(par))
}