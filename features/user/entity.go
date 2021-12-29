package user

type UserCore struct{
	Id					uint
	Nik					string
	Name				string
	PhoneNumber			string
	Email				string
	Password			string
	Token				string
}

type Service interface{
	RegisterUser(data UserCore)(err error)
	LoginUser(data UserCore)(user UserCore, err error)
	GetUsers(data UserCore)(users []UserCore, err error)
	GetUserById(id int) (user UserCore, err error)
	UpdateUser(data UserCore)error
}

type Repository interface{
	InsertUserData(data UserCore)(id int, err error)
	CheckUser(data UserCore)(user UserCore, err error)
	GetData(UserCore)(user []UserCore, err error)
	GetDataById(id int)(user UserCore, err error)
	UpdateUser(data UserCore)error
	GetUserByEmail(email string) (bool, error)
	GetUserByNik(nik string)(bool, error)
}