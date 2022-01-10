package participant

import "time"

type ParticipantCore struct {
	ID          uint
	Nik         uint
	Fullname    string
	Address     string
	PhoneNumber string
	UserID      uint
	VacID       uint
	Status      string
	AppliedAt   time.Time
	Vac         VacCore
	User        UserCore
}

type UserCore struct {
	ID          uint
	Nik         string
	Name        string
	PhoneNumber string
	Email       string
}

type VacCore struct {
	ID          int
	Description string
	Location    string
	Latitude    float64
	Longitude   float64
	Sessions    []SessionCore
	VacType     string
	Stock       int
	AdminId     int
}

type SessionCore struct {
	ID          uint
	VacId       uint
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type Repository interface {
	ApplyParticipant(ParticipantCore) error
	GetParticipantByUserID(int) ([]ParticipantCore, error)
	GetParticipantByID(int) (ParticipantCore, error)
	GetParticipantByVacID(int) ([]ParticipantCore, error)
	GetParticipantMultiParam(int, int) (ParticipantCore, error)
	RejectParticipant(int) error
	AcceptParticipant(int) error
}

type Service interface {
	ApplyParticipant(ParticipantCore) error
	GetParticipantByUserID(int) ([]ParticipantCore, error)
	GetParticipantByID(int) (ParticipantCore, error)
	GetParticipantByVacID(int) ([]ParticipantCore, error)
	GetParticipantMultiParam(int, int) (ParticipantCore, error)
	RejectParticipant(int, int) error
	AcceptParticipant(int, int) error
}
