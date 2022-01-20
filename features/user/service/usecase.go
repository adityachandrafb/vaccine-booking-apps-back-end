package service

import (
	"errors"
	"fmt"
	"vac/features/user"
	"vac/helper"
	"vac/middleware"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) user.Service {
	return &userService{userRepository}
}

func (us *userService) RegisterUser(data user.UserCore) error {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) || !helper.ValidateNik(data.Nik) || !helper.ValidatePhoneNumber(data.PhoneNumber) || len(data.Name) == 0 {
		return errors.New("data kurang lengkap")
	}

	isExist, err := us.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		return err
	}
	if isExist {
		msg := fmt.Sprintf("email %v tidak bisa didaftarkan", data.Email)
		return errors.New(msg)
	}

	isExistNIK, err := us.userRepository.GetUserByNik(data.Nik)
	if err != nil {
		return err
	}
	if isExistNIK {
		msg := fmt.Sprintf("nik %v tidak bisa didaftarkan", data.Nik)
		return errors.New(msg)
	}

	userId, err := us.userRepository.InsertUserData(data)
	if err != nil {
		return err
	}
	if userId < 0 {
		return err
	}
	return nil
}

func (us *userService) GetUsers(data user.UserCore) ([]user.UserCore, error) {
	users, err := us.userRepository.GetData(data)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userService) LoginUser(data user.UserCore) (user.UserCore, error) {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) {
		return user.UserCore{}, errors.New("email atau password salah")
	}
	userData, err := us.userRepository.CheckUser(data)
	if err != nil {
		return user.UserCore{}, err
	}

	userData.Token, err = middleware.CreateToken(userData.Id, "user")
	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (us *userService) GetUserById(id int) (user.UserCore, error) {
	userData, err := us.userRepository.GetDataById(id)

	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (us *userService) GetUserOwnIdentity(id int) (user.UserCore, error) {
	userData, err := us.userRepository.GetDataById(id)

	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (us *userService) UpdateUser(data user.UserCore) error {
	if data.Nik != "" {
		if !helper.ValidateNik(data.Nik) {
			return errors.New("data kurang lengkap")
		}

		isExistNIK, err := us.userRepository.GetUserByNik(data.Nik)
		if err != nil {
			return err
		}
		if isExistNIK {
			msg := fmt.Sprintf("nik %v tidak bisa didaftarkan", data.Nik)
			return errors.New(msg)
		}
	}

	if data.Email != "" {
		if !helper.ValidateEmail(data.Email) {
			return errors.New("data kurang lengkap")
		}

		isExist, err := us.userRepository.GetUserByEmail(data.Email)
		if err != nil {
			return err
		}
		if isExist {
			msg := fmt.Sprintf("email %v tidak bisa didaftarkan", data.Email)
			return errors.New(msg)
		}
	}

	if data.PhoneNumber != "" {
		if !helper.ValidatePhoneNumber(data.PhoneNumber) {
			return errors.New("incomplete or invalid data")
		}
	}

	if data.Password != "" {
		if !helper.ValidatePassword(data.Password) {
			return errors.New("data kurang lengkap")
		}
	}

	err := us.userRepository.UpdateUser(data)
	if err != nil {
		return err
	}

	return nil
}
