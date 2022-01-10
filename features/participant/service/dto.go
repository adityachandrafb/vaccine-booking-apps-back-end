package service

import (
	"vac/features/participant"
	"vac/features/user"
	"vac/features/vac"
)

func ToUserCore(data user.UserCore) participant.UserCore {
	return participant.UserCore{
		ID:          data.Id,
		Nik:         data.Nik,
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
	}
}

func ToVacCore(data vac.VacCore) participant.VacCore {
	return participant.VacCore{
		ID:          data.ID,
		Description: data.Description,
		Location:    data.Location,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
		Sessions:    ToSessionCore(data.Sessions),
		VacType:     data.VacType,
		Stock:       data.Stock,
		AdminId:     data.AdminId,
	}
}

func ToSessionCore(data []vac.SessionCore) []participant.SessionCore {
	converted := []participant.SessionCore{}
	for _, ses := range data {
		converted = append(converted, participant.SessionCore{
			ID:          ses.ID,
			Description: ses.Description,
			StartTime:   ses.StartTime,
			EndTime:     ses.EndTime})
	}
	return converted
}
