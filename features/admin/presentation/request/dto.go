package request

import "vac/features/admin"

type AdminRequest struct {
	Name     string `json: "name"`
	Position string `json: "position"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

type AdminLogin struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func FromAdminRequest(data AdminRequest) admin.AdminCore {
	return admin.AdminCore{
		Name: data.Name,
		Position: data.Position,
		Email: data.Email,
		Password: data.Password,
	}
}

func FromAdminLogin(data AdminLogin) admin.AdminCore {
	return admin.AdminCore{
		Email: data.Email,
		Password: data.Password,
	}
}