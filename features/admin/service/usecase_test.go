package service

import (
	"errors"
	"os"
	"testing"
	"vac/features/admin"
	"vac/features/admin/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	adminRepo  mocks.Repository
	adminServ  admin.Service
	adminData  admin.AdminCore
	adminLogin admin.AdminCore
	adminsData []admin.AdminCore
)

func TestMain(m *testing.M) {
	adminServ = NewAdminService(&adminRepo)
	adminData = admin.AdminCore{
		ID:			1,
		Name:		"Rumah Sakit Universitas Sumatera Utara",
		Position:	"xxx",
		Email:		"rs.usu@gmail.com",
		Password:	"123456",
	}
	adminLogin = admin.AdminCore{
		Email:		"rs.usu@gmail.com",
		Password:	"123456",
	}
	adminsData = []admin.AdminCore{
		{
			ID:      	1,
			Name:		"Rumah Sakit Universitas Sumatera Utara",
			Position:	"xxx",
			Email:		"rs.usu@gmail.com",
			Password:	"123456",
		},
	}
	os.Exit(m.Run())
}

func TestRegisterAdmin(t *testing.T) {
	t.Run("Register admin success", func(t *testing.T) {
		adminRepo.On("CreateAdmin", mock.AnythingOfType("admin.AdminCore")).Return(nil).Once()
		err := adminServ.RegisterAdmin(adminData)
		assert.Nil(t, err)
	})

	t.Run("Register admin invalid data", func(t *testing.T) {
		err := adminServ.RegisterAdmin(admin.AdminCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "incomplete or invalid data", err.Error())
	})

	t.Run("Register admin invalid data", func(t *testing.T) {
		adminRepo.On("CreateAdmin", mock.AnythingOfType("admin.AdminCore")).Return(errors.New("error create admin")).Once()
		err := adminServ.RegisterAdmin(adminData)
		assert.NotNil(t, err)
		assert.Equal(t, "error create admin", err.Error())
	})

}

func TestLoginAdmin(t *testing.T) {
	t.Run("Login admin success", func(t *testing.T) {
		adminRepo.On("CheckAdmin", mock.AnythingOfType("admin.AdminCore")).Return(adminData, nil).Once()
		resp, err := adminServ.LoginAdmin(adminLogin)
		assert.Nil(t, err)
		assert.Equal(t, adminData.Email, resp.Email)
		assert.Equal(t, adminData.Name, resp.Name) //responnya apa ya?
	})

	t.Run("Login admin error invalid data", func(t *testing.T) {
		resp, err := adminServ.LoginAdmin(admin.AdminCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid data", err.Error())
		assert.Equal(t, "", resp.Name)
		assert.Equal(t, "", resp.Email)
	})

	t.Run("Login admin error check admin", func(t *testing.T) {
		adminRepo.On("CheckAdmin", mock.AnythingOfType("admin.AdminCore")).Return(admin.AdminCore{}, errors.New("error check admin")).Once()
		resp, err := adminServ.LoginAdmin(adminLogin)
		assert.NotNil(t, err)
		assert.Equal(t, "error check admin", err.Error())
		assert.Equal(t, "", resp.Name) //responnya apa
		assert.Equal(t, "", resp.Email)
	})
}

func TestGetAdmins(t *testing.T) {
	t.Run("get admins success", func(t *testing.T) {
		adminRepo.On("GetAdmins").Return(adminsData, nil).Once()
		resp, err := adminServ.GetAdmins()
		assert.Nil(t, err)
		assert.Equal(t, len(adminsData), len(resp))
		assert.Equal(t, adminsData[0].Name, resp[0].Name) //respon name
	})

	t.Run("get admins error get admins", func(t *testing.T) {
		adminRepo.On("GetAdmins").Return(nil, errors.New("error get admins")).Once()
		resp, err := adminServ.GetAdmins()
		assert.NotNil(t, err)
		assert.Equal(t, "error get admins", err.Error())
		assert.Nil(t, resp)
	})
}

func TestGetAdminById(t *testing.T) {
	t.Run("get admin by id success", func(t *testing.T) {
		adminRepo.On("GetAdminById", mock.AnythingOfType("admin.AdminCore")).Return(adminData, nil).Once()
		resp, err := adminServ.GetAdminById(admin.AdminCore{ID: 1})
		assert.Nil(t, err)
		assert.Equal(t, adminData.Email, resp.Email)
		assert.Equal(t, adminData.Name, resp.Name) //respon Name
	})

	t.Run("get admin by id error GetAdminById", func(t *testing.T) {
		adminRepo.On("GetAdminById", mock.AnythingOfType("admin.AdminCore")).Return(admin.AdminCore{}, errors.New("error get admin by id")).Once()
		resp, err := adminServ.GetAdminById(admin.AdminCore{ID: 1})
		assert.NotNil(t, err)
		assert.Equal(t, "error get admin by id", err.Error())
		assert.Equal(t, "", resp.Email)
		assert.Equal(t, "", resp.Name) //respon name
	})
}
