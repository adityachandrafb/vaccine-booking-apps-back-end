package request

import (
	"vac/features/user"

)

type UserRequest struct {
	ID          int   `json:"id"`
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UserAuth struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (data *UserAuth) ToUserCore() user.UserCore{
	return user.UserCore{
		Email: data.Email,
		Password: data.Password,
	}
}

func (requestdata *UserRequest) ToUserCore() user.UserCore{
	return user.UserCore{
		Id: uint(requestdata.ID),
		Nik: requestdata.Nik,
		Name: requestdata.Name,
		PhoneNumber: requestdata.PhoneNumber,
		Email: requestdata.Email,
		Password: requestdata.Password,
	}
}