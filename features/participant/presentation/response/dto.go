package response

import (
	"net/http"
	"time"
	"vac/features/participant"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
	Data    interface{}
}

func NewErrorResponse(e echo.Context, msg string, code int) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}

func NewSuccessResponse(e echo.Context, msg string, data interface{}) error {
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data:    data,
	})
}

type ParticipantResponse struct {
	ID          uint
	Nik         string
	Fullname    string
	Address     string
	PhoneNumber string
	UserID      uint
	VacID       uint
	SessionID   uint
	Status      string
	Vac         VacResponse
	User        UserResponse
	Session     SessionResponse
}

type ParticipantResponseUser struct {
	ID          uint
	Nik         string
	Fullname    string
	Address     string
	PhoneNumber string
	UserID      uint
	VacID       uint
	SessionID   uint
	Status      string
	Vac         VacResponse
	Session     SessionResponse
}

type ParticipantResponseVac struct {
	ID          uint
	Nik         string
	Fullname    string
	Address     string
	PhoneNumber string
	UserID      uint
	VacID       uint
	SessionID   uint
	Status      string
	User        UserResponse
	Session     SessionResponse
}

type VacResponse struct {
	ID          uint
	Description string
	Location    string
	Address		string
	// Sessions    []SessionResponse
	VacType string
	Stock   int
}

type VacDetailResponse struct {
	ID          int
	Description string
	Location    string
	Address		string
	Latitude    float64
	Longitude   float64
	// Sessions    []SessionResponse
	VacType string
	Stock   int
	AdminId int
}

type SessionResponse struct {
	ID          uint
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type UserResponse struct {
	ID          uint
	Name        string
	PhoneNumber string
	Email       string
}

type UserDetailResponse struct {
	ID          uint
	Nik         string
	Name        string
	PhoneNumber string
	Email       string
}

func ToParticipantResponseUser(data participant.ParticipantCore) ParticipantResponseUser {
	return ParticipantResponseUser{
		ID:          data.ID,
		Nik:         data.Nik,
		Fullname:    data.Fullname,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		UserID:      data.UserID,
		VacID:       data.VacID,
		SessionID:   data.SessionID,
		Status:      data.Status,
		Vac:         ToVacResponse(data.Vac),
		Session:     ToSessionsResponse(data.Sessions),
	}
}

func ToVacResponse(data participant.VacCore) VacResponse {
	return VacResponse{
		ID:          uint(data.ID),
		Description: data.Description,
		Location:    data.Location,
		Address: data.Address,
		VacType:     data.VacType,
		Stock:       data.Stock,
	}
}

func ToUserResponse(data participant.UserCore) UserResponse {
	return UserResponse{
		ID:          data.ID,
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
	}
}

func ToParticipantResponseVac(data participant.ParticipantCore) ParticipantResponseVac {
	return ParticipantResponseVac{
		ID:          data.ID,
		Nik:         data.Nik,
		Fullname:    data.Fullname,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		UserID:      data.UserID,
		VacID:       data.VacID,
		SessionID:   data.SessionID,
		Status:      data.Status,
		User:        ToUserResponse(data.User),
		Session:     ToSessionsResponse(data.Sessions),
	}
}

func ToParticipantResponseVacList(data []participant.ParticipantCore) []ParticipantResponseVac {
	convertedData := []ParticipantResponseVac{}
	for _, par := range data {
		convertedData = append(convertedData, ToParticipantResponseVac(par))
	}
	return convertedData
}

func ToParticipantResponseUserList(data []participant.ParticipantCore) []ParticipantResponseUser {
	convertedData := []ParticipantResponseUser{}
	for _, par := range data {
		convertedData = append(convertedData, ToParticipantResponseUser(par))
	}
	return convertedData
}

func ToParticipantResponse(data participant.ParticipantCore) ParticipantResponse {
	return ParticipantResponse{
		ID:          data.ID,
		Nik:         data.Nik,
		Fullname:    data.Fullname,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		UserID:      data.UserID,
		VacID:       data.VacID,
		SessionID:   data.SessionID,
		Status:      data.Status,
		Vac:         ToVacResponse(data.Vac),
		User:        ToUserResponse(data.User),
		Session:     ToSessionsResponse(data.Sessions),
	}
}

func ToVacDetailResponse(data participant.VacCore) VacDetailResponse {
	return VacDetailResponse{
		ID:          data.ID,
		Description: data.Description,
		Location:    data.Location,
		Address: data.Address,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
		VacType:     data.VacType,
		Stock:       data.Stock,
		AdminId:     data.AdminId,
	}
}

func ToSessionsResponse(data participant.SessionCore) SessionResponse {
	// converted := []SessionResponse{}
	// for _, ses := range data {
	// 	temp := SessionResponse{
	// 		ses.ID,
	// 		ses.Description,
	// 		ses.StartTime,
	// 		ses.EndTime,
	// 	}
	// 	converted = append(converted, temp)
	// }
	// return converted
	return SessionResponse{
		ID:          data.ID,
		Description: data.Description,
		StartTime:   data.StartTime,
		EndTime:     data.EndTime,
	}
}

func ToUserDetailResponse(data participant.UserCore) UserDetailResponse {
	return UserDetailResponse{
		ID:          data.ID,
		Nik:         data.Nik,
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
	}
}
