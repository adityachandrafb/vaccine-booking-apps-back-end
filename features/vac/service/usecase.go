package service

import (
	"errors"
	"fmt"
	"strconv"
	"vac/features/vac"
	"vac/helper"
)

type vacUseCase struct {
	vacRepository vac.Repository
}

func NewVacUseCase(vacRepository vac.Repository) vac.Service {
	return &vacUseCase{vacRepository}
}

func (vu *vacUseCase) CreateVaccinationPost(data vac.VacCore) error {
	stockString := strconv.Itoa(data.Stock)
	if helper.IsEmpty(data.Description) || helper.IsEmpty(data.Location) || helper.IsEmpty(stockString) {
		return errors.New("Invalid data")
	}
	err := vu.vacRepository.InsertData(data)
	if err != nil {
		return err
	}
	return nil
}

func (vu *vacUseCase) GetVaccinationPost(data vac.VacCore) ([]vac.VacCore, error) {
	vacData, err := vu.vacRepository.GetVacData(data)
	if err != nil {
		return nil, err
	}
	return vacData, err
}

func (vu *vacUseCase) GetVaccinationByIdPost(id int) (vac.VacCore, error) {
	vacData, err := vu.vacRepository.GetVacDataById(id)
	if err != nil {
		return vac.VacCore{}, nil
	}
	return vacData, nil
}

func (vu *vacUseCase) DeleteVaccinationPost(data vac.VacCore) (err error) {
	vacData, err := vu.vacRepository.GetVacDataById(data.ID)
	if err != nil {
		return err
	}
	if vacData.AdminId != data.AdminId {
		msg := fmt.Sprintf("admin with id %v does not have vaccination post with id %v", data.AdminId, data.ID)
		return errors.New(msg)
	}
	err = vu.vacRepository.DeleteVacData(data)
	if err != nil {
		return err
	}
	return nil
}

func (vu *vacUseCase) UpdateVaccinationPost(data vac.VacCore) error {
	stockString := strconv.Itoa(data.Stock)
	if helper.IsEmpty(data.Description) || helper.IsEmpty(data.Location) || helper.IsEmpty(stockString) {
		return errors.New("Invalid data")
	}
	vacData, err := vu.vacRepository.GetVacDataById(data.ID)
	if err != nil {
		return err
	}
	if vacData.AdminId != data.AdminId {
		msg := fmt.Sprintf("admin with id %v does not have vaccination post with id %v", data.AdminId, data.ID)
		return errors.New(msg)
	}
	err = vu.vacRepository.UpdateVacData(data)
	if err != nil {
		return err
	}
	return err

}
