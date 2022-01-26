package service

import (
	"errors"
	"os"
	"testing"
	"time"
	"vac/features/vac"
	"vac/features/vac/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	vacService    vac.Service
	vacRepository mocks.Repository
	vacData       vac.VacCore
	vacsData      []vac.VacCore
)

func TestMain(m *testing.M) {
	vacService = NewVacUseCase(&vacRepository)
	vacData = vac.VacCore{
		ID:				1,
		Description:	"Sinovac Dosis Pertama",
		Location:		"RS Universitas Sumatra Utara",
		Latitude: 		10,
		Longitude: 		10,
		VacType:		"Sinovac",
		Stock:			200,
		AdminId	:		1,
		Sessions: 		[]vac.SessionCore{
			{
				Description: 	"Dosis Pertama tanggal 13 Januari",
				StartTime:		time.Now(),
				EndTime: 		time.Now().Add(time.Hour * 2),
			},
			{
				Description: 	"Dosis Pertama tanggal 13 Januari",
				StartTime:		time.Now(),
				EndTime: 		time.Now().Add(time.Hour * 2),
			},
		},
	}
	vacsData = []vac.VacCore{
		{
			ID:      		1,
			Description:	"Sinovac Dosis Pertama",
			Location:		"RS Universitas Sumatra Utara",
			Latitude: 		10,
			Longitude: 		10,
			VacType:		"Sinovac",
			Stock:			200,
			AdminId	:		1,
			Sessions: 		[]vac.SessionCore{
				{
					Description: 	"Dosis Pertama tanggal 13 Januari",
					StartTime:		time.Now(),
					EndTime: 		time.Now().Add(time.Hour * 1),
				},
				{
					Description: 	"Dosis Pertama tanggal 13 Januari",
					StartTime:		time.Now(),
					EndTime: 		time.Now().Add(time.Hour * 1),
				},
			},
		},
	}
	os.Exit(m.Run())
}

func TestCreateVaccinationPost(t *testing.T) {
	t.Run("create vaccination success", func(t *testing.T) {
		vacRepository.On("InsertData", mock.AnythingOfType("vac.VacCore")).Return(nil).Once()
		err := vacService.CreateVaccinationPost(vacData)
		assert.NotNil(t, err)
	})

	t.Run("create vaccination failed invalid data", func(t *testing.T) {
		err := vacService.CreateVaccinationPost(vac.VacCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "deskripsi harus ada", err.Error())
	})

	t.Run("create vaccination failed InsertData", func(t *testing.T) {
		vacRepository.On("InsertData", mock.AnythingOfType("vac.VacCore")).Return(errors.New("error insert data")).Once()
		err := vacService.CreateVaccinationPost(vacData)
		assert.NotNil(t, err)
		assert.Equal(t, "alamat harus ada", err.Error())
	})
}

func TestGetVaccinationPost(t *testing.T) {
	t.Run("get vaccination post success", func(t *testing.T) {
		vacRepository.On("GetVacData", mock.AnythingOfType("vac.VacCore")).Return(vacsData, nil).Once()
		resp, err := vacService.GetVaccinationPost(vacData)
		assert.Nil(t, err)
		assert.Equal(t, len(vacsData), len(resp))
		assert.Equal(t, vacsData[0].VacType, resp[0].VacType) //VacType atau Description
	})

	t.Run("get vaccination post error GetVaccinationData", func(t *testing.T) {
		vacRepository.On("GetVacData", mock.AnythingOfType("vac.VacCore")).Return(nil, errors.New("error get vaccination data")).Once()
		resp, err := vacService.GetVaccinationPost(vacData)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.Equal(t, "error get vaccination data", err.Error())
	})
}

func TestGetVaccinationByIdPost (t *testing.T) {
	t.Run("get vaccination by id success", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		resp, err := vacService.GetVaccinationByIdPost(1)
		assert.Nil(t, err)
		assert.Equal(t, vacData.VacType, resp.VacType) //VacType atau Description
		assert.Equal(t, len(vacData.Sessions), len(resp.Sessions))
	})

	t.Run("get vaccination post by id error GetVaccinationByIdPost", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vac.VacCore{}, errors.New("error get vaccination data by id")).Once()
		resp, err := vacService.GetVaccinationByIdPost(1)
		assert.Nil(t, err)
		assert.Equal(t, "", resp.VacType) //VacType atau Description
		assert.Equal(t, 0, len(resp.Sessions))
	})
}

func TestDeleteVaccinationPost(t *testing.T) {
	t.Run("delete vaccination post success", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		vacRepository.On("DeleteVacData", mock.AnythingOfType("vac.VacCore")).Return(nil).Once()
		err := vacService.DeleteVaccinationPost(vacData)
		assert.Nil(t, err)
	})

	t.Run("delete vaccination post error get job data by id", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vac.VacCore{}, errors.New("error get vaccination data by id")).Once()
		err := vacService.DeleteVaccinationPost(vacData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get vaccination data by id", err.Error())
	})

	t.Run("delete vaccination post error DeleteVacData", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		vacRepository.On("DeleteVacData", mock.AnythingOfType("vac.VacCore")).Return(errors.New("error delete vaccination data")).Once()
		err := vacService.DeleteVaccinationPost(vacData)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete vaccination data", err.Error())
	})

}

func TestUpdateVaccinationPost(t *testing.T) {
	t.Run("update vaccination success", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		vacRepository.On("UpdateVacData", mock.AnythingOfType("vac.VacCore")).Return(nil).Once()
		err := vacService.UpdateVaccinationPost(vacData)
		assert.Nil(t, err)
	})

	t.Run("update vaccination error GetVacDataById", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vac.VacCore{}, errors.New("error get vaccination data by id")).Once()
		err := vacService.UpdateVaccinationPost(vacData)
		assert.NotNil(t, err)
		assert.Equal(t, "error get vaccination data by id", err.Error())
	})

	t.Run("update vaccination error UpdateVacData", func(t *testing.T) {
		vacRepository.On("GetVacDataById", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		vacRepository.On("UpdateVacData", mock.AnythingOfType("vac.VacCore")).Return(errors.New("error update vaccination data")).Once()
		err := vacService.UpdateVaccinationPost(vacData)
		assert.NotNil(t, err)
		assert.Equal(t, "error update vaccination data", err.Error())
	})

	t.Run("update vaccination error invalid data", func(t *testing.T) {
		err := vacService.UpdateVaccinationPost(vac.VacCore{})
		assert.NotNil(t, err)
		assert.Equal(t, "invalid data. make sure description, location, and stock available", err.Error())
	})
}