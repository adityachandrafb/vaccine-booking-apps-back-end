package driver

import(
	UserModel "vac/features/user/data"
	AdminModel "vac/features/admin/data"
)
type Entity struct{
	Model interface{}
}

func registerEntities() []Entity{
	return []Entity{
		{UserModel.User{}},
		{AdminModel.Admin{}},
	}
}