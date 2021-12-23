package user

type UserCore struct{
	Id					uint
	Email				string
	Fullname			string
	PhoneNumber			string
	Password			string
	Token				string
	ParticipantVacReg	[]ParticipantVacRegCore
}

type ParticipantVacRegCore struct{
	Id					uint
	Nama				string
	NIK					string
	NomorTelepon		string
	Alamat				string
}