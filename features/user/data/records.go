package data

import (
	"gorm.io/gorm"
	"vac/features/user"
)
type User struct {
	gorm.Model
	// Id					uint
	Nik					string
	Name				string
	PhoneNumber			string
	Email				string
	Password			string
}

func toUserRecord(user user.UserCore)User{
	return User{
		Model: gorm.Model{ID: user.Id},
		Nik: user.Nik,
		Name: user.Name,
		PhoneNumber: user.PhoneNumber,
		Email: user.Email,
		Password: user.Password,
	}
}
func toUserCore(u User)user.UserCore{
	return user.UserCore{
		Id: u.ID,
		Nik: u.Nik,
		Name: u.Name,
		PhoneNumber: u.PhoneNumber,
		Email: u.Email,
		Password: u.Password,
	}
}

func toUserCoreList(uList []User) []user.UserCore{
	convertedUser:=[]user.UserCore{}
	for _,user:=range uList{
		convertedUser=append(convertedUser, toUserCore(user))
	}
	return convertedUser
}