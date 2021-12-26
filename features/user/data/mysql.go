package data

import (
	"errors"
	"vac/features/user"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB)user.Repository{
	return &mysqlUserRepository{DB}
}

func (mr *mysqlUserRepository)InsertUserData(data user.UserCore)(int,error){
	recordData:=toUserRecord(data)
	err:=mr.DB.Create(&recordData).Error
	if err!=nil{
		return 0, err
	}
	return int(recordData.ID),nil
}

func (mr *mysqlUserRepository)CheckUser(data user.UserCore)(user.UserCore, error){
	var userData User
	err:=mr.DB.Where("email=? and password = ?", data.Email, data.Password).First(&userData).Error

	if userData.Name=="" && userData.ID==0{
		return user.UserCore{}, errors.New("no existing user")
	}
	if err!=nil{
		return user.UserCore{}, err
	}
	return toUserCore(userData), nil
}


func (mr *mysqlUserRepository)GetDataById(id int)(user.UserCore, error){
	var userData User
	err:=mr.DB.First(&userData, id).Error

	if userData.Name == "" && userData.ID == 0 {
		return user.UserCore{}, errors.New("no existing user")
	}
	if err != nil {
		return user.UserCore{}, err
	}

	return toUserCore(userData), nil
}


func (mr *mysqlUserRepository) GetData(data user.UserCore) ([]user.UserCore, error) {
	var users []User
	err:=mr.DB.Find(&users).Error

	if err!=nil{
		return nil, err
	}
	return toUserCoreList(users), nil
}

func (mr *mysqlUserRepository)UpdateUser(data user.UserCore)error{
	err:=mr.DB.Debug().Model(&User{}).Where("id=?",data.Id).Updates(User{
		Nik: data.Nik,
		Name:data.Name,
		PhoneNumber: data.PhoneNumber,
		Email: data.Email,
		Password: data.Password,
	}).Error
	if err!=nil{
		return nil
	}
	return nil
}

func (mr *mysqlUserRepository) GetUserByEmail(email string) (bool, error) {
	var userModel User
	err := mr.DB.Where("email = ?", email).Find(&userModel).Error
	if err != nil {
		return false, err
	}
	if userModel.ID != 0 {
		return true, nil
	}
	return false, nil
}