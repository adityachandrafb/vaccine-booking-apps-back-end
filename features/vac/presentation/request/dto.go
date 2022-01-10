package request

import (
	"time"
	"vac/features/vac"
)

type Vac struct {
	ID          uint		`json:"id"`
	Description string		`json:"description"`
	Location    string		`json:"location"`
	Latitude    float64		`json:"latitude"`
	Longitude   float64		`json:"longitude"`
	Sessions    []Session	`json:"sessions"`
	VacType     string		`json:"vacType"`
	Stock       int			`json:"stock"`
	AdminId     int			`json:"adminId"`
}

type VacUpdate struct {
	ID          uint		`json:"id"`
	Description string		`json:"description"`
	Location    string		`json:"location"`
	Latitude    float64		`json:"latitude"`
	Longitude   float64		`json:"longitude"`
	Sessions    []Session	`json:"sessions"`
	VacType     string		`json:"vacType"`
	Stock       int			`json:"stock"`
	AdminId     int			`json:"adminId"`
}

type Session struct {
	ID          uint		`json:"id"`
	Description string		`json:"description"`
	StartTime   time.Time	`json:"startTime"`
	EndTime     time.Time	`json:"endTime"`
}

type VacFilter struct{
	Description string		`json:"description"`
	Location    string		`json:"location"`
	Latitude    float64		`json:"latitude"`
	Longitude   float64		`json:"longitude"`
}

func (v *Vac) ToCore() vac.VacCore{
	convertedSession:=[]vac.SessionCore{}
	for _,req:=range v.Sessions{
		convertedSession = append(convertedSession, vac.SessionCore{
			ID: req.ID,
			Description: req.Description,
			StartTime: req.StartTime,
			EndTime: req.EndTime,
		})
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

func (vf *VacFilter) ToCore() vac.VacCore{
	return vac.VacCore{
		Description: vf.Description,
		Location: vf.Location,
		Latitude: vf.Latitude,
		Longitude: vf.Longitude,
	}
}

func (s *Session) ToCore() vac.SessionCore{
	return vac.SessionCore{
		ID: s.ID,
		Description: s.Description,
		StartTime: s.StartTime,
		EndTime: s.EndTime,
	}
}


func (vu *VacUpdate) ToCore() vac.VacCore{
	convertedSession:=[]vac.SessionCore{}
	for _,req:=range vu.Sessions{
		convertedSession = append(convertedSession, req.ToCore())
	}
	return vac.VacCore{
		ID: int(vu.ID),
		Description: vu.Description,
		Location: vu.Location,
		Latitude: vu.Latitude,
		Longitude: vu.Longitude,
		Sessions: convertedSession,
		VacType: vu.VacType,
		Stock: vu.Stock,
		AdminId: vu.AdminId,
	}
}