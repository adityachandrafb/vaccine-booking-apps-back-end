package response

import (
	"net/http"
	"vac/features/admin"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

type AdminLoginResponse struct {
	ID	     uint	`json: "id"`
	Name     string `json: "name"`
	Position string `json: "position"`
	Token	 string `json:	"token"`
}

type AdminResponse struct {
	ID	     uint	`json: "id"`
	Name     string `json: "name"`
	Position string `json: "position"`
}

func NewSuccessResponse(e echo.Context, msg string, data interface{})error{
	return e.JSON(http.StatusOK, Response{
		Message: msg,
		Data: data,
	})
}
func NewErrorResponse(e echo.Context, msg string, code int)error{
	return e.JSON(http.StatusOK, Response{
		Message: msg,
	})
}

func ToAdminLoginResponse(data admin.AdminCore)AdminLoginResponse{
	return AdminLoginResponse{
		ID: data.ID,
		Name: data.Name,
		Position: data.Position,
		Token: data.Token,
	}
}

func ToAdminResponse(data admin.AdminCore)AdminResponse{
	return AdminResponse{
		ID: data.ID,
		Name: data.Name,
		Position: data.Position,
	}
}

func ToAdminResponseList(data []admin.AdminCore)[]AdminResponse{
	convertedRec:=[]AdminResponse{}
	for _, rec:=range data{
		convertedRec = append(convertedRec, ToAdminResponse(rec))
	}
	return convertedRec
}