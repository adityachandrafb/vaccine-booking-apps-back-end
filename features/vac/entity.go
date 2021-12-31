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
	CreateVaccination(data VacCore)(err error)
	GetVaccination(data VacCore)([]VacCore, error)
	GetVaccinationById(id int)(VacCore, error)
	DeleteVaccination(data VacCore)(err error)
	UpdateVaccination(data VacCore)error
}

type Repository interface{
	InsertData(data VacCore)(err error)
	GetVacData(data VacCore)([]VacCore, error)
	GetVacDataById(id int)(VacCore, error)
	DeleteVacData(data VacCore)error
	UpdateVacData(data VacCore)error
}