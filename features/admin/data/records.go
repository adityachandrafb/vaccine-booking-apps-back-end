package data

import (
	"vac/features/admin"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name		string
	Position	string
	Email		string
	Password	string
}

func FromCore(data admin.AdminCore) Admin{
	return Admin{
		Name: data.Name,
		Position: data.Position,
		Email: data.Email,
		Password: data.Password,
	}
}

func ToCore(data Admin) admin.AdminCore{
	return admin.AdminCore{
		ID: data.ID,
		Name: data.Name,
		Position: data.Position,
		Email: data.Email,
		Password: data.Password,
	}
}

func ToCoreList(data []Admin) []admin.AdminCore{
	convertedData:=[]admin.AdminCore{}
	for _, rec:=range data{
		convertedData=append(convertedData, ToCore(rec))
	}
	return convertedData
}