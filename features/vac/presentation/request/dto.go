package request

import (
	"time"
	"vac/features/vac"
)

type Vac struct {
	ID          uint		`json:"id"`
	Description string		`json:"description"`
	Location    string		`json:"location"`
	Address		string		`json:"address"`
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
	Address		string		`json:"address"`
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
	StartTime   string		`json:"startTime"`
	EndTime     string		`json:"endTime"`
}

type VacFilter struct{
	Description string		`json:"description"`
	Location    string		`json:"location"`
	Address		string		`json:"address"`
	Latitude    float64		`json:"latitude"`
	Longitude   float64		`json:"longitude"`
	Radius		float64		`json:"radius"`
}

func (v *Vac) ToCore() vac.VacCore{
	layoutTime := "2006-01-02T15:04"
	convertedSession:=[]vac.SessionCore{}
	for _,req:=range v.Sessions{

		starttimeconv,_:=time.Parse(layoutTime, req.StartTime)
		endtimeconv,_:=time.Parse(layoutTime, req.EndTime)
		convertedSession = append(convertedSession, vac.SessionCore{
			ID: req.ID,
			Description: req.Description ,
			StartTime: starttimeconv,
			EndTime: endtimeconv,
		})
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

func (vf *VacFilter) ToCore() vac.VacCore{
	return vac.VacCore{
		Description: vf.Description,
		Location: vf.Location,
		Address: vf.Address,
		Latitude: vf.Latitude,
		Longitude: vf.Longitude,
	}
}

func (s *Session) ToCore() vac.SessionCore{
	layoutTime := "2006-01-02T15:04"
	starttimeconv,_:=time.Parse(layoutTime, s.StartTime)
	endtimeconv,_:=time.Parse(layoutTime, s.EndTime)
	return vac.SessionCore{
		ID: s.ID,
		Description: s.Description,
		StartTime: starttimeconv,
		EndTime: endtimeconv,
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
		Address: vu.Address,
		Latitude: vu.Latitude,
		Longitude: vu.Longitude,
		Sessions: convertedSession,
		VacType: vu.VacType,
		Stock: vu.Stock,
		AdminId: vu.AdminId,
	}
}