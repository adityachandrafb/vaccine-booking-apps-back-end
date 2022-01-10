package driver

import (
	AdminModel "vac/features/admin/data"
	UserModel "vac/features/user/data"
	VacModel "vac/features/vac/data"
	ParticipantModel "vac/features/participant/data"
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
		{ParticipantModel.Participant{}},
	}
}