package request

import "vac/features/participant"

type ParticipantRequest struct {
	Nik         uint   `json:"nik"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	VacID       uint   `json:"VacID"`
	UserID      uint   `json:"UserID`
}

func (par *ParticipantRequest) ToCore() participant.ParticipantCore{
	return participant.ParticipantCore{
		Nik: par.Nik,
		Fullname: par.Fullname,
		Address: par.Address,
		PhoneNumber: par.PhoneNumber,
		VacID: par.VacID,
		UserID: par.UserID,
	}
}