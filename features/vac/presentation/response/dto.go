package response

import (
	"net/http"
	"time"
	"vac/features/vac"

	"github.com/labstack/echo"
)

type Response struct {
	Message string
	Data    interface{}
}

type VacResponse struct {
	ID          uint
	Description string
	Location    string
	Address		string
	Latitude    float64
	Longitude   float64
	Sessions    []SessionResponse
	VacType     string
	Stock       int
	AdminId     int
}

type SessionResponse struct {
	ID          int
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

func NewSuccessResponse(e echo.Context, msg string, data interface{})error{
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data: data,
	})
}

func NewErrorResponse(e echo.Context, msg string, code int)error{
	return e.JSON(code, Response{
		Message: msg,
	})
}

func ToVacResponse(data vac.VacCore)VacResponse{
	convertedSessions:=[]SessionResponse{}
	for _, ses:=range data.Sessions{
		convertedSessions=append(convertedSessions, SessionResponse{
			ID: int(ses.ID),
			Description: ses.Description,
			StartTime: ses.StartTime,
			EndTime: ses.EndTime,
		})
	}
	return VacResponse{
		ID: uint(data.ID),
		Description: data.Description,
		Location: data.Location,
		Address: data.Address,
		Latitude: data.Latitude,
		Longitude: data.Longitude,
		Sessions: convertedSessions,
		VacType: data.VacType,
		Stock: data.Stock,
		AdminId: data.AdminId,
	}
}

func ToVacResponseList(data []vac.VacCore) []VacResponse {
	convertedVac:=[]VacResponse{}
	for _, vac := range data {
		convertedVac = append(convertedVac, ToVacResponse(vac))
	}
	return convertedVac
}