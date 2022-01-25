package data

import (
	"errors"
	"vac/features/participant"

	"gorm.io/gorm"
)

type mysqlParRepository struct {
	DB *gorm.DB
}

func NewMysqlParticipantRepository(db *gorm.DB) participant.Repository {
	return &mysqlParRepository{db}
}

func (pr *mysqlParRepository) UpdateParticipant(data participant.ParticipantCore) error {
	parData:= ToParticipantRecord(data)
	err:=pr.DB.Debug().Where("id = ?", data.ID).Updates(&parData).Error
	if err!=nil{
		return err
	}
	return nil
}
func (pr *mysqlParRepository) DeleteParticipant(data participant.ParticipantCore) error {
	err:=pr.DB.Debug().Delete(&Participant{}, data.ID).Error
	if err!=nil{
		return err
	}
	return nil
}
func (pr *mysqlParRepository) CountParicipantByUserId(userId int) (int, error) {
	var countPar int64
	var parModel Participant
	err := pr.DB.Model(&parModel).Where("user_id=?", userId).Count(&countPar).Error
	if err != nil {
		return 0, err
	}
	return int(countPar), nil
}

func (pr *mysqlParRepository) CountParticipantByVac(vacId int) (int, error) {
	var countPar int64
	var parModel Participant
	err := pr.DB.Model(&parModel).Where("vac_id = ?", vacId).Count(&countPar).Error
	if err != nil {
		return 0, err
	}
	return int(countPar), nil
}

func (pr *mysqlParRepository) GetParticipantByNIK(nik string) (bool, error) {
	var parModel Participant
	err := pr.DB.Where("nik = ?", nik).Find(&parModel).Error
	if err != nil {
		return false, err
	}
	if parModel.ID != 0 {
		return true, nil
	}
	return false, nil
}

func (pr *mysqlParRepository) ApplyParticipant(data participant.ParticipantCore) error {
	parData := ToParticipantRecord(data)
	err := pr.DB.Create(&parData).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *mysqlParRepository) GetParticipantByUserID(id int) ([]participant.ParticipantCore, error) {
	var participants []Participant
	err := pr.DB.Debug().Where("user_id=?", id).Joins("Vac").Joins("Session").Find(&participants).Error
	if err != nil {
		return nil, err
	}
	return ToCoreList(participants), err
}

func (pr *mysqlParRepository) RejectParticipant(id int) error {
	err := pr.DB.Model(&Participant{}).Where("id=?", id).Update("status", "canceled").Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *mysqlParRepository) AcceptParticipant(id int) error {
	err := pr.DB.Model(&Participant{}).Where("id=?", id).Update("status", "VACCINATED").Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *mysqlParRepository) GetParticipantByID(id int) (participant.ParticipantCore, error) {
	var data Participant
	err := pr.DB.Debug().Joins("Session").First(&data, id).Error
	if err != nil {
		return participant.ParticipantCore{}, err
	}
	if data.ID == 0 {
		return participant.ParticipantCore{}, errors.New("record not found")
	}
	return ToCore(data), nil
}

func (pr *mysqlParRepository) GetParticipantByVacID(id int) ([]participant.ParticipantCore, error) {
	var participants []Participant
	err := pr.DB.Where("vac_id=?", id).Preload("Session").Joins("User").Find(&participants).Error
	if err != nil {
		return nil, err
	}
	return ToCoreList(participants), err
}

func (pr *mysqlParRepository) GetParticipantMultiParam(vacId int, userId int) (participant.ParticipantCore, error) {
	var data Participant
	err := pr.DB.Where("vac_id=? and user_id=?", vacId, userId).Find(&data).Error
	if err != nil {
		return participant.ParticipantCore{}, err
	}
	return ToCore(data), nil
}
