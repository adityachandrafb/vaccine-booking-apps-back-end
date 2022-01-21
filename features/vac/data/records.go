package data

import (
	"time"
	"vac/features/vac"

	"gorm.io/gorm"
)

type Vac struct {
    ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`

	Description	string
	Location	string
	Address		string
	Latitude 	float64
	Longitude 	float64
	Distance	float64	
	Sessions	[]Session
	VacType		string
	Stock		int
	AdminId		int
}

type Session struct{	
    ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`

	VacId		uint
	Description	string
	StartTime 	time.Time
	EndTime		time.Time
}

func toRecordSession(req vac.SessionCore)Session{
	return Session{
		ID: req.ID,
		VacId: req.VacId,
		Description: req.Description,
		StartTime: req.StartTime,
		EndTime: req.EndTime,
	}
}

func toRecordVac(data vac.VacCore)Vac{
	convertedSession :=[]Session{}
	for _,req:=range data.Sessions{
		convertedSession = append(convertedSession, toRecordSession(req))
	}
	return Vac{
		ID: uint(data.ID),
		Description	: data.Description, 
		Location	: data.Location,
		Address		: data.Address,
		Latitude 	: data.Latitude,
		Longitude 	: data.Longitude,
		Sessions	: convertedSession,
		VacType		: data.VacType,
		Stock		: data.Stock,
		AdminId		: data.AdminId,
		
	}
}

func (v *Vac) toCore() vac.VacCore{
	convertedSession:=[]vac.SessionCore{}
	for _, req:=range v.Sessions{
		convertedSession = append(convertedSession, req.toCore())
	}
	return vac.VacCore{
		ID: int(v.ID),
		Description: v.Description,
		Location: v.Location,
		Address: v.Address,
		Latitude: v.Latitude,
		Longitude: v.Longitude,
		Sessions: convertedSession,
		VacType: v.VacType,
		Stock: v.Stock,
		AdminId: v.AdminId,

	}
}

func (s *Session) toCore() vac.SessionCore{
	return vac.SessionCore{
		ID: s.ID,
		VacId: s.VacId,
		Description: s.Description,
		StartTime: s.StartTime,
		EndTime: s.EndTime,

	}
}

func toCoreList(vacs []Vac) []vac.VacCore{
	var convertedData []vac.VacCore
	for _, vac:=range vacs{
		convertedData=append(convertedData, vac.toCore())
	}
	return convertedData
}

func SeparateVacSession(data Vac)(Vac, []Session){
	newVac:= Vac{
		ID: data.ID,
		Description: data.Description,
		Location: data.Location,
		Address: data.Address,
		Latitude: data.Latitude,
		Longitude: data.Longitude,
		VacType: data.VacType,
		Stock: data.Stock,
		AdminId: data.AdminId,
	}
	newSessions:=data.Sessions
	
	return newVac, newSessions
}