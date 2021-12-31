package data

import (
	"time"
	"vac/features/vac"

	"gorm.io/gorm"
)

type Vac struct {
	gorm.Model
	Description	string
	Location	string
	Latitude 	float64
	Longitude 	float64
	Sessions	[]Session
	VacType		string
	Stock		int
	AdminId		int
}

type Session struct{	
	gorm.Model
	VacId		uint
	Description	string
	StartTime 	time.Time
	EndTime		time.Time
}

func toRecordSession(req vac.SessionCore)Session{
	return Session{
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
		Description	: data.Description, 
		Location	: data.Location,
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
		Description: data.Description,
		Location: data.Location,
		Latitude: data.Latitude,
		Longitude: data.Longitude,
		VacType: data.VacType,
		Stock: data.Stock,
		AdminId: data.AdminId,
	}
	newSessions:=data.Sessions
	return newVac, newSessions
}