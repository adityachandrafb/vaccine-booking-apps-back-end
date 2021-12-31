package vac

import "time"

type VacCore struct {
	ID       	int
	Description	string
	Location	string
	Latitude 	float64
	Longitude 	float64
	Sessions	[]SessionCore
	VacType		string
	Stock		int
	AdminId		int
}

type SessionCore struct{
	ID			uint
	VacId		uint
	Description	string
	StartTime 	time.Time
	EndTime		time.Time
}

type Service interface{
	CreateVaccinationPost(data VacCore)(err error)
	GetVaccinationPost(data VacCore)([]VacCore, error)
	GetVaccinationByIdPost(id int)(VacCore, error)
	DeleteVaccinationPost(data VacCore)(err error)
	UpdateVaccinationPost(data VacCore)error
}

type Repository interface{
	InsertData(data VacCore)(err error)
	GetVacData(data VacCore)([]VacCore, error)
	GetVacDataById(id int)(VacCore, error)
	DeleteVacData(data VacCore)error
	UpdateVacData(data VacCore)error
}