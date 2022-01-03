package driver

import(
	UserModel "vac/features/user/data"
	AdminModel "vac/features/admin/data"
	VacModel "vac/features/vac/data"
)
type Entity struct{
	Model interface{}
}

func registerEntities() []Entity{
	return []Entity{
		{UserModel.User{}},
		{AdminModel.Admin{}},
		{VacModel.Vac{}},
		{VacModel.Session{}},
	}
}