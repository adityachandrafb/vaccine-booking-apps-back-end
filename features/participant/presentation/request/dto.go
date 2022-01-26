package request

import "vac/features/participant"

type ParticipantRequest struct {
	Nik         string `json:"nik"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	VacID       uint   `json:"VacID"`
	SessionID   uint   `json:"session_id"`
	UserID      uint   `json:"UserID`
}
type ParticipantUpdateRequest struct {
	ID          uint	`json:"id"`
	Nik         string `json:"nik"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	VacID       uint   `json:"VacID"`
	SessionID   uint   `json:"session_id"`
	UserID      uint   `json:"UserID`
}

func (par *ParticipantRequest) ToCore() participant.ParticipantCore {
	return participant.ParticipantCore{
		Nik:         par.Nik,
		Fullname:    par.Fullname,
		Address:     par.Address,
		PhoneNumber: par.PhoneNumber,
		VacID:       par.VacID,
		SessionID:   par.SessionID,
		UserID:      par.UserID,
	}
}

func (par *ParticipantUpdateRequest) ToCore() participant.ParticipantCore {
	return participant.ParticipantCore{
		ID: par.ID,
		Nik:         par.Nik,
		Fullname:    par.Fullname,
		Address:     par.Address,
		PhoneNumber: par.PhoneNumber,
		VacID:       par.VacID,
		SessionID:   par.SessionID,
		UserID:      par.UserID,
	}
}
