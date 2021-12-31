package request

import (
	"time"
	"vac/features/vac"
)

type Vac struct {
	ID          int			`json:"id"`
	Description string		`json:"description"`
	Location    string		`json:"location"`
	Latitude    float64		`json:"latitude"`
	Longitude   float64		`json:"longitude"`
	Sessions    []string	`json:"sessions"`
	VacType     string		`json:"vacType"`
	Stock       int			`json:"stock"`
	AdminId     int			`json:"adminId"`
}
type VacUpdate struct {
	ID          int			`json:"id"`
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
	ID          uint
	Description string
	StartTime   time.Time
	EndTime     time.Time
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
			Description: req,
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
	convertedSessions:=[]vac.SessionCore{}
	for _, ses:=range vu.Sessions{
		convertedSessions = append(convertedSessions, ses.ToCore())
	}
	return vac.VacCore{
		ID: int(vu.ID),
		Description: vu.Description,
		Location: vu.Location,
		Latitude: vu.Latitude,
		Longitude: vu.Longitude,
		Sessions: convertedSessions,
		VacType: vu.VacType,
		Stock: vu.Stock,
		AdminId: vu.AdminId,
	}
}