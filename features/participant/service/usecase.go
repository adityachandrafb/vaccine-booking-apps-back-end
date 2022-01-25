package service

import (
	"errors"
	"fmt"
	"time"
	"vac/features/participant"
	"vac/features/user"
	"vac/features/vac"
	"vac/helper"
)

type parService struct {
	parRepository participant.Repository
	vacService    vac.Service
	userService   user.Service
}

func NewParService(pr participant.Repository, vs vac.Service, us user.Service) participant.Service {
	return &parService{
		parRepository: pr,
		vacService:    vs,
		userService:   us,
	}
}

func (pr *parService) DeleteParticipant(data participant.ParticipantCore) error {
	parData, err:=pr.parRepository.GetParticipantByID(int(data.ID))
	if err!=nil{
		return err
	}
	if parData.UserID!=data.UserID{
		msg := fmt.Sprintf("user with id %v does not have vaccination post with id %v", data.UserID, data.ID)
		return errors.New(msg)
	}
	err=pr.parRepository.DeleteParticipant(data)
	if err!=nil{
		return err
	}
	return nil
}



func (pr *parService) ApplyParticipant(data participant.ParticipantCore) error {
	vacData, err := pr.vacService.GetVaccinationByIdPost(int(data.VacID))
	if err != nil {
		return err
	}

	if vacData.ID == 0 {
		msg := fmt.Sprintf("vac with id %v not found", vacData.ID)
		return errors.New(msg)
	}

	stockUsed, err := pr.parRepository.CountParticipantByVac(int(data.VacID))
	if err != nil {
		msg := fmt.Sprintf("vac with id %v failed to count", vacData.ID)
		return errors.New(msg)
	}
	if stockUsed > vacData.Stock+1 {
		msg := fmt.Sprintf("Mohon maaf vaksin di %v sudah habis. Boleh dicoba di tempat lain ya 🤓", vacData.Location)
		return errors.New(msg)
	}

	limitUser, err := pr.parRepository.CountParicipantByUserId(int(data.UserID))
	if err != nil {
		msg := fmt.Sprintf("vac with id %v failed to count", vacData.ID)
		return errors.New(msg)
	}
	if limitUser > 5 {
		msg := "mohon maaf, kamu sudah melebihi batas untuk mendaftarkan partisipan 😭"
		return errors.New(msg)
	}

	if !helper.ValidateNik(data.Nik) {
		return errors.New("panjang nik harus 16 karakter 😡")
	}

	if  !helper.ValidatePhoneNumber(data.PhoneNumber) {
		return errors.New("nomor telepon hanya boleh 8-15 angka 👌🏻")
	}

	if len(data.Fullname) == 0 {
		return errors.New("nama partisipan jangan lupa 😇")
	}

	if  len(data.Address) == 0 {
		return errors.New("alamat partisipan diisi ya 😻")
	}

	if data.SessionID == 0 {
		return errors.New("sesinya jangan lupa diisi ya syg 😘")
	}

	isExist, err := pr.parRepository.GetParticipantByNIK(data.Nik)
	if err != nil {
		return err
	}
	if isExist {
		msg := "nik yang kamu masukkan sudah terdaftar 🤭"
		return errors.New(msg)
	}

	if data.Status == "" {
		data.Status = "registered"
	}

	data.AppliedAt = time.Now()
	err = pr.parRepository.ApplyParticipant(data)
	if err != nil {
		return nil
	}
	return nil
}

func (pr *parService) GetParticipantByUserID(id int) ([]participant.ParticipantCore, error) {
	participants, err := pr.parRepository.GetParticipantByUserID(id)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (pr *parService) GetParticipantByID(id int) (participant.ParticipantCore, error) {
	parData, err := pr.parRepository.GetParticipantByID(id)
	if err != nil {
		return participant.ParticipantCore{}, err
	}
	userData, err := pr.userService.GetUserById(int(parData.UserID))
	if err != nil {
		return participant.ParticipantCore{}, err
	}
	vacData, err := pr.vacService.GetVaccinationByIdPost(int(parData.VacID))
	if err != nil {
		return participant.ParticipantCore{}, err
	}
	parData.User = ToUserCore(userData)
	parData.Vac = ToVacCore(vacData)

	return parData, nil
}

func (pr *parService) GetParticipantByVacID(id int) ([]participant.ParticipantCore, error) {
	data, err := pr.parRepository.GetParticipantByVacID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (pr *parService) GetParticipantMultiParam(id int, userId int) (participant.ParticipantCore, error) {
	data, err := pr.parRepository.GetParticipantMultiParam(id, userId)
	if err != nil {
		return participant.ParticipantCore{}, err
	}
	return data, nil
}

func (pr *parService) RejectParticipant(id int, adminId int) error {
	data, err := pr.GetParticipantByID(id)
	if err != nil {
		msg := fmt.Sprintf("participant with id %v not found", id)
		return errors.New(msg)
	}
	if data.Vac.AdminId != adminId {
		msg := fmt.Sprintf("admin with id %v not allowed to access post with id %v", adminId, id)
		return errors.New(msg)
	}
	if data.Status != "registered" {
		msg := fmt.Sprintf("this participant has been %v", data.Status)
		return errors.New(msg)
	}
	err = pr.parRepository.RejectParticipant(id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *parService) AcceptParticipant(id int, adminId int) error {
	data, err := pr.GetParticipantByID(id)
	if err != nil {
		msg := fmt.Sprintf("participant with id %v not found", id)
		return errors.New(msg)
	}
	if data.Vac.AdminId != adminId {
		msg := fmt.Sprintf("admin with id %v not allowed to access post with id %v", adminId, id)
		return errors.New(msg)
	}
	if data.Status != "registered" {
		msg := fmt.Sprintf("this participant has been %v", data.Status)
		return errors.New(msg)
	}
	err = pr.parRepository.AcceptParticipant(id)
	if err != nil {
		return err
	}
	return nil
}
