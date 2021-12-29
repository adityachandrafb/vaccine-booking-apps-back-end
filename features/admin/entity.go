package admin

type AdminCore struct{
	ID			uint
	Name		string
	Position	string
	Email		string
	Password	string
	Token		string
}

type Service interface{
	RegisterAdmin(data AdminCore)error
	LoginAdmin(data AdminCore)(AdminCore, error)
	GetAdmins()([]AdminCore, error)
	GetAdminById(data AdminCore)(AdminCore, error)
}

type Repository interface{
	CreateAdmin(data AdminCore)error
	CheckAdmin(data AdminCore)(AdminCore, error)
	GetAdmins()([]AdminCore, error)
	GetAdminById(data AdminCore)(AdminCore, error)
}