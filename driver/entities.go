package driver

import(
	UserModel "vac/features/user/data"
)
type Entity struct{
	Model interface{}
}

func registerEntities() []Entity{
	return []Entity{
		{UserModel.User{}},
	}
}