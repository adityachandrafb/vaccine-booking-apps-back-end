package data

import (
	"vac/features/vac"

	"gorm.io/gorm"
)

type mysqlVaccineRepository struct {
	DB *gorm.DB
}

func NewMysqlVaccineRepository(DB *gorm.DB)vac.Repository{
	return &mysqlVaccineRepository{DB}
}

func (vr *mysqlVaccineRepository)InsertData(data vac.VacCore)error{
	recordData:=toRecordVac(data)
	result:=vr.DB.Create(&recordData)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

func (vr *mysqlVaccineRepository)GetVacData(data vac.VacCore)([]vac.VacCore, error){
	var vacs []Vac
	err:=vr.DB.Debug().Joins("JOIN admins ON vacs.admin_id=admins.id").Preload("Sessions").Find(&vacs).Error
	if err!=nil{
		return nil, err
	}
	return toCoreList(vacs), nil
}

func (vr *mysqlVaccineRepository)GetVacDataById(id int)(vac.VacCore, error){
	var vacData Vac
	err:=vr.DB.Preload("Sessions").First(&vacData, id).Error
	if err!=nil{
		return vac.VacCore{}, err
	}
	return vacData.toCore(), nil
}

func (vr *mysqlVaccineRepository)DeleteVacData(data vac.VacCore)error{
	err:=vr.DB.Debug().Delete(&Vac{}, data.ID).Error
	if err!=nil{
		return err
	}

	err=vr.DB.Debug().Where("vac_id=?",data.ID).Delete(&Session{}).Error
	if err!=nil{
		return err
	}
	return nil
}

func (vr *mysqlVaccineRepository)UpdateVacData(data vac.VacCore)error{
	vacData, sessions:=SeparateVacSession(toRecordVac(data))

	err:=vr.DB.Debug().Where("id = ?",data.ID).Updates(&vacData).Error
	if err!=nil{
		return err
	}

	for _,req:=range sessions{
		if req.ID!=0{
			if req.Description==""{
				err=vr.DB.Debug().Delete(&Session{}, req.ID).Error
				if err!=nil{
					return err
				}
			}else{
				err=vr.DB.Debug().Model(&Session{}).Where("id=?",req.ID).Update("description", req.Description).Error
				if err!=nil{
					return err
				}
			}
		}else if req.ID==0{
			req.VacId=uint(data.ID)
			err=vr.DB.Debug().Select("VacID", "Description").Create(&req).Error
		}
		if err!=nil{
			return err
		}
	}
	return nil
}