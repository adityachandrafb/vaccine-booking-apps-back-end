package data

import (
	"errors"
	"vac/features/admin"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) admin.Repository{
	return &AdminRepository{DB}
}

func (ar *AdminRepository)CreateAdmin(data admin.AdminCore)error{
	convertedData:=FromCore(data)
	err:=ar.DB.Create(&convertedData).Error
	if err!=nil{
		return err
	}
	return nil
}

func (ar *AdminRepository)CheckAdmin(data admin.AdminCore)(admin.AdminCore, error){
	var adminData Admin

	err:=ar.DB.Where("email = ? and password = ?", data.Email, data.Password).First(&adminData).Error
	if err!=nil{
		return admin.AdminCore{}, err
	}
	if adminData.ID==0 && adminData.Email==""{
		return admin.AdminCore{}, errors.New("No data exist for admin")
	}

	return ToCore(adminData), err
}
func (ar *AdminRepository)GetAdmins()([]admin.AdminCore, error){
	var admins []Admin
	err:=ar.DB.Find(&admins).Error
	if err!=nil{
		return nil, err
	}

	return ToCoreList(admins),nil
}

func (ar *AdminRepository)GetAdminById(data admin.AdminCore)(admin.AdminCore, error){
	var adminData Admin
	err:= ar.DB.First(&adminData, data.ID).Error
	if err!=nil{
		return admin.AdminCore{}, err
	}
	return ToCore(adminData), nil
}