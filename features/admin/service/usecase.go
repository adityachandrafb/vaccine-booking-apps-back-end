package service

import (
	"errors"
	"vac/features/admin"
	"vac/helper"
	"vac/middleware"
)

type adminService struct {
	adminRepository admin.Repository
}

func NewAdminService(adminRepo admin.Repository)admin.Service{
	return &adminService{adminRepo}
}

func (as *adminService)RegisterAdmin(data admin.AdminCore)error{
	if !helper.ValidateEmail(data.Email)||!helper.ValidatePassword(data.Password){
		return errors.New("incomplete or invalid data")
	}
	err:=as.adminRepository.CreateAdmin(data)
	if err!=nil{
		return err
	}
	return nil
}

func (as *adminService) LoginAdmin(data admin.AdminCore) (admin.AdminCore, error){
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password){
		return admin.AdminCore{}, errors.New("invalid data")
	}
	data, err:=as.adminRepository.CheckAdmin(data)
	if err!=nil{
		return admin.AdminCore{}, err
	}
	
	data.Token, err = middleware.CreateToken(data.ID, "admin")
	if err!=nil{
		return admin.AdminCore{}, err
	}
	return data, nil
}

func (as *adminService)GetAdmins()([]admin.AdminCore,error){
	admins, err:=as.adminRepository.GetAdmins()
	if err!=nil{
		return nil, err
	}
	return admins, nil
}

func (as *adminService)GetAdminById(data admin.AdminCore)(admin.AdminCore, error){
	adminData, err:=as.adminRepository.GetAdminById(data)
	if err!=nil{
		return admin.AdminCore{}, err
	}
	return adminData, nil
}

