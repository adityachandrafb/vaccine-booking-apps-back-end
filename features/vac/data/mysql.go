package data

import (
	// "fmt"
	"vac/features/vac"

	"gorm.io/gorm"
)

type mysqlVaccineRepository struct {
	DB *gorm.DB
}

func NewMysqlVaccineRepository(DB *gorm.DB) vac.Repository {
	return &mysqlVaccineRepository{DB}
}

func (vr *mysqlVaccineRepository) GetNearbyFacilities(latitude float64, longitude float64) ([]vac.VacCore, error) {
	var vacs []Vac
	err:=vr.DB.Debug().Where("latitude BETWEEN ? and ? and longitude between ? and ?", latitude-0.1, latitude+0.1, longitude-0.1, longitude+0.1).Find(&vacs).Error
	if err!=nil{
		return nil, err
	}
	return toCoreList(vacs), nil
}

func (vr *mysqlVaccineRepository) InsertData(data vac.VacCore) error {
	recordData := toRecordVac(data)
	result := vr.DB.Create(&recordData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (vr *mysqlVaccineRepository) GetVacData(data vac.VacCore) ([]vac.VacCore, error) {
	var vacs []Vac
	err := vr.DB.Debug().Joins("JOIN admins ON vacs.admin_id=admins.id").Preload("Sessions").Find(&vacs).Error
	if err != nil {
		return nil, err
	}
	return toCoreList(vacs), nil
}

func (vr *mysqlVaccineRepository) GetVacDataById(id int) (vac.VacCore, error) {
	var vacData Vac
	err := vr.DB.Preload("Sessions").First(&vacData, id).Error
	if err != nil {
		return vac.VacCore{}, err
	}
	return vacData.toCore(), nil
}

func (vr *mysqlVaccineRepository) DeleteVacData(data vac.VacCore) error {
	err := vr.DB.Debug().Delete(&Vac{}, data.ID).Error
	if err != nil {
		return err
	}

	err = vr.DB.Debug().Where("vac_id=?", data.ID).Delete(&Session{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (vr *mysqlVaccineRepository) UpdateVacData(data vac.VacCore) error {
	vacData, sessions := SeparateVacSession(toRecordVac(data))

	err := vr.DB.Debug().Where("id = ?", data.ID).Updates(&vacData).Error
	if err != nil {
		return err
	}
	// fmt.Printf("%+v\n",vacData)
	for _, ses := range sessions {
		// fmt.Printf("%+v\n",ses)
		if ses.ID != 0 {
			if ses.Description == "" {
				err = vr.DB.Debug().Delete(&Session{}, ses.ID).Error
				if err != nil {
					return err
				}
			} else {
				err = vr.DB.Debug().Model(&Session{}).Where("id = ?", ses.ID).Updates(ses).Error
				if err != nil {
					return err
				}
			}
		} else if ses.ID == 0 {
			ses.VacId = vacData.ID
			err = vr.DB.Debug().Create(&ses).Error
		}
		if err != nil {
			return err
		}
	}
	return nil
}
