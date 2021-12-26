package response

import "vac/features/user"

type UserResponse struct {
	Id          uint   `json:"id"`
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
}

type UserLoginResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToUserResponse(user user.UserCore)UserResponse{
	return UserResponse{
		Id: user.Id,
		Nik: user.Nik,
		Name: user.Name,
		PhoneNumber: user.PhoneNumber,
		Email: user.Email,
	}
}


func ToUserResponseList(userList []user.UserCore) []UserResponse {
	convertedUser := []UserResponse{}
	for _, user := range userList {
		convertedUser = append(convertedUser, ToUserResponse(user))
	}

	return convertedUser
}

func TouserLoginResponse(user user.UserCore)UserLoginResponse{
	return UserLoginResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}