package service

import (
	"errors"
	"os"
	"testing"
	"vac/features/user"
	"vac/features/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo   mocks.Repository
	userServ   user.Service
	usersData  []user.UserCore
	userData   user.UserCore
	userLogin  user.UserCore
	userUpdate user.UserCore
)

func TestMain(m *testing.M) {
	userServ = NewUserService(&userRepo)

	usersData = []user.UserCore{
		{
			Id:				1,
			Nik:			"3309074302000004",
			Name:       	"Aditya Chandra",
			PhoneNumber: 	"081234567891",
			Email:      	"aditya@gmail.com",
			Password:   	"aditya123",
		},
	}
	
	userData = user.UserCore{
		Nik:			"3309074302000004",
		Name:       	"Aditya Chandra",
		PhoneNumber: 	"081234567891",
		Email:      	"aditya@gmail.com",
		Password:   	"aditya123",
	}

	userLogin = user.UserCore{
		Email:    "aditya@gmail.com",
		Password: "aditya123",
	}

	userUpdate = user.UserCore{
		Name: 	"Aditya test",
		// Email:	"Adityatest@gmail.com",	
			
	}
	os.Exit(m.Run())
}

func TestGetUser(t *testing.T) {
	t.Run("validate get users", func(t *testing.T) {
		userRepo.On("GetData", mock.AnythingOfType("user.UserCore")).Return(usersData, nil).Once()
		resp, err := userServ.GetUsers(user.UserCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get users", func(t *testing.T) {
		userRepo.On("GetData", mock.AnythingOfType("user.UserCore")).Return(nil, errors.New("error on db"))
		resp, err := userServ.GetUsers(user.UserCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestRegisterUser(t *testing.T) {
	t.Run("Register user success", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("GetUserByNik", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(1, nil).Once()

		err := userServ.RegisterUser(userData)
		assert.Nil(t, err)
	})

	//EMAIL
	t.Run("Register user error invalid email", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := userServ.RegisterUser(user.UserCore{
			Email: "023ujawol",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "data kurang lengkap")
	})

	t.Run("Register user error GetUserByEmail", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})

	t.Run("Register user error email exist", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(true, nil).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})

	t.Run("Register user error insert data", func(t *testing.T) {
		userRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepo.On("InsertUserData", mock.AnythingOfType("user.UserCore")).Return(0, errors.New("error")).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})


	//NIK
	t.Run("Register user error invalid nik", func(t *testing.T) {
		userRepo.On("GetUserByNik", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := userServ.RegisterUser(user.UserCore{
			Nik: "12345678",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "data kurang lengkap")
	})

	t.Run("Register user error nik exist", func(t *testing.T) {
		userRepo.On("GetUserByNik", mock.AnythingOfType("string")).Return(true, nil).Once()
		err := userServ.RegisterUser(userData)
		assert.NotNil(t, err)
	})

}

func TestLoginUser(t *testing.T) {
	t.Run("Login user success", func(t *testing.T) {
		userRepo.On("CheckUser", mock.AnythingOfType("user.UserCore")).Return(userData, nil).Once()
		data, err := userServ.LoginUser(userLogin)
		assert.Equal(t, userData.Email, data.Email)
		assert.Nil(t, err)
	})

	t.Run("Login failed email invalid", func(t *testing.T) {
		data, err := userServ.LoginUser(user.UserCore{
			Email:    "9012uhjja",
			Password: "fran123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "email atau password salah")
		assert.Empty(t, data.Email)
	})

	t.Run("Login error check user", func(t *testing.T) {
		userRepo.On("CheckUser", mock.AnythingOfType("user.UserCore")).Return(user.UserCore{}, errors.New("error check data")).Once()
		data, err := userServ.LoginUser(userLogin)
		assert.Equal(t, "error check data", err.Error())
		assert.NotNil(t, err)
		assert.Empty(t, data.Id)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("Get user by id success", func(t *testing.T) {
		userRepo.On("GetDataById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		data, err := userServ.GetUserById(1)
		assert.Equal(t, userData.Id, data.Id)
		assert.Nil(t, err)
	})

	t.Run("Get user by id error", func(t *testing.T) {
		userRepo.On("GetDataById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error get user")).Once()
		data, err := userServ.GetUserById(1)
		assert.Empty(t, data)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
	})
}

func TestGetUserByOwnIdentity(t *testing.T) {
	t.Run("Get user by own identity success", func(t *testing.T) {
		userRepo.On("GetDataById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		data, err := userServ.GetUserById(1)
		assert.Equal(t, userData.Id, data.Id)
		assert.Nil(t, err)
	})

	t.Run("Get user by own identity error", func(t *testing.T) {
		userRepo.On("GetDataById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error get user")).Once()
		data, err := userServ.GetUserById(1)
		assert.Empty(t, data)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Update user success", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.AnythingOfType("user.UserCore")).Return(nil).Once()

		err := userServ.UpdateUser(userUpdate)
		assert.Nil(t, err)
	})

	t.Run("Update user error UpdateUserData", func(t *testing.T) {
		userRepo.On("UpdateUser", mock.AnythingOfType("user.UserCore")).Return(errors.New("error update user")).Once()
		err := userServ.UpdateUser(userUpdate)
		assert.NotNil(t, err)
		assert.Equal(t, "error update user", err.Error())
	})

	//VALIDATE NIK
	t.Run("Update failed nik invalid", func(t *testing.T) {
		err := userServ.UpdateUser(user.UserCore{
			Nik:    "12345678",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "data kurang lengkap")
	})

	//VALIDATE EMAIL dan PASSWORD
	t.Run("Update failed email invalid", func(t *testing.T) {
		err := userServ.UpdateUser(user.UserCore{
			Email:    "9012uhjja",
			Password: "fran123",
			PhoneNumber:    "12345678888",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "data kurang lengkap")
	})

	
	t.Run("Login failed email invalid", func(t *testing.T) {
		data, err := userServ.LoginUser(user.UserCore{
			Email:    "9012uhjja",
			Password: "fran123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "email atau password salah")
		assert.Empty(t, data.Email)
	})

	
}
