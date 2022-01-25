package participant

import "time"

type ParticipantCore struct {
	ID          uint
	Nik         string
	Fullname    string
	Address     string
	PhoneNumber string
	UserID      uint
	VacID       uint
	SessionID   uint
	Status      string
	AppliedAt   time.Time
	Vac         VacCore
	User        UserCore
	Sessions    SessionCore
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
	Address		string
	Latitude    float64
	Longitude   float64
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
	//validation
	GetParticipantByNIK(nik string) (bool, error)
	CountParticipantByVac(vacId int) (int, error)
	CountParicipantByUserId(userId int) (int, error)
	DeleteParticipant(ParticipantCore)error
	UpdateParticipant(data ParticipantCore)error
}

type Service interface {
	ApplyParticipant(ParticipantCore) error
	GetParticipantByUserID(int) ([]ParticipantCore, error)
	GetParticipantByID(int) (ParticipantCore, error)
	GetParticipantByVacID(int) ([]ParticipantCore, error)
	GetParticipantMultiParam(int, int) (ParticipantCore, error)
	RejectParticipant(int, int) error
	AcceptParticipant(int, int) error
	DeleteParticipant(ParticipantCore) error
	UpdateParticipant(data ParticipantCore)error
}
