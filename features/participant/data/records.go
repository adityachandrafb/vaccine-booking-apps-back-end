package data

import (
	"time"
	"vac/features/participant"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	Nik         uint
	Fullname    string
	Address     string
	PhoneNumber string
	UserID      uint
	VacID       uint
	Status      string
	AppliedAt   time.Time
	Vac         Vac
	User        User
}

type Vac struct {
	gorm.Model
	Description string
	Location    string
	Latitude    float64
	Longitude   float64
	Sessions    []Session
	VacType     string
	Stock       int
	AdminId     int
}

type Session struct {
	gorm.Model
	VacId       uint
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type User struct {
	gorm.Model
	Nik         string
	Name        string
	PhoneNumber string
	Email       string
}

func (v *Vac) toCore() participant.VacCore {
	convertedSession := []participant.SessionCore{}
	for _, ses := range v.Sessions {
		convertedSession = append(convertedSession, ses.toCore())
	}
	return participant.VacCore{
		ID:          int(v.ID),
		Description: v.Description,
		Location:    v.Location,
		Latitude:    v.Latitude,
		Longitude:   v.Longitude,
		Sessions:    convertedSession,
		VacType:     v.VacType,
		Stock:       v.Stock,
		AdminId:     v.AdminId,
	}
}

func (u User) toCore() participant.UserCore {
	return participant.UserCore{
		ID:          u.ID,
		Nik:         u.Nik,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
}

func (s *Session) toCore() participant.SessionCore {
	return participant.SessionCore{
		ID:          s.ID,
		VacId:       s.VacId,
		Description: s.Description,
		StartTime:   s.StartTime,
		EndTime:     s.EndTime,
	}
}

func toCoreList(vacs []Vac) []participant.VacCore {
	var convertedData []participant.VacCore
	for _, vac := range vacs {
		convertedData = append(convertedData, vac.toCore())
	}
	return convertedData
}

func ToParticipantRecord(data participant.ParticipantCore) Participant {
	return Participant{

		Nik:         data.Nik,
		Fullname:    data.Fullname,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		UserID:      data.UserID,
		VacID:       data.VacID,
		Status:      data.Status,
		AppliedAt:   data.AppliedAt,
	}
}

func ToCore(data Participant) participant.ParticipantCore {
	return participant.ParticipantCore{
		ID:          data.ID,
		Nik:         data.Nik,
		Fullname:    data.Fullname,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		UserID:      data.UserID,
		VacID:       data.VacID,
		Status:      data.Status,
		AppliedAt:   data.AppliedAt,
		Vac:         data.Vac.toCore(),
		User:        data.User.toCore(),
	}
}

func ToCoreList(data []Participant) []participant.ParticipantCore {
	convertedData := []participant.ParticipantCore{}
	for _, par := range data {
		convertedData = append(convertedData, ToCore(par))
	}
	return convertedData
}
