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
	err := pr.DB.Debug().Where("user_id=?", id).Joins("Vac").Preload("Vac.Sessions").Find(&participants).Error
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
	err := pr.DB.Debug().First(&data, id).Error
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
	err := pr.DB.Where("vac_id=?", id).Joins("User").Find(&participants).Error
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
