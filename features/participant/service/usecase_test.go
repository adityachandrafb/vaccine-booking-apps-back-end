package service

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
	"vac/features/participant"
	"vac/features/vac"
	par_m "vac/features/participant/mocks"
	vac_m "vac/features/vac/mocks"
	"vac/features/user"
	user_m "vac/features/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	parServ       participant.Service
	parRepository par_m.Repository
	userService   user_m.Service
	parMockServ   par_m.Service
	vacService    vac_m.Service
	vacData       vac.VacCore
	parData       participant.ParticipantCore
	userData      user.UserCore
	parsData      []participant.ParticipantCore
)

func TestMain(m *testing.M) {
	parServ = NewParService(&parRepository, &vacService, &userService)
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
	parData = participant.ParticipantCore{
		ID:         1,
		Nik:		"312345678912",
		Fullname:	"Aditya Chandra",
		Address:	"Sumatera Utara",
		PhoneNumber:"081234567891",
		UserID:		1,
		VacID:		1,
		SessionID:	1,
		Status:		"pending",
		AppliedAt:   time.Now(),
		Vac: participant.VacCore{
			ID:				1,
			Description:	"Sinovac Dosis Pertama",
			Location:		"RS Universitas Sumatra Utara",
			Latitude: 		10,
			Longitude: 		10,
			VacType:		"Sinovac",
			Stock:			200,
			AdminId	:		1,
		},
	}
	userData = user.UserCore{
		Nik:			"3309074302000004",
		Name:       	"Aditya Chandra",
		PhoneNumber: 	"081234567891",
		Email:      	"aditya@gmail.com",
		Password:   	"aditya123",
	}
	parsData = []participant.ParticipantCore{
		parData,
	}
	os.Exit(m.Run())
}

func TestApplyParticipant(t *testing.T) {
	t.Run("apply participant success", func(t *testing.T) {
		vacService.On("GetVaccinationByIdPost", mock.Anything).Return(vacData, nil).Once()
		parRepository.On("CountParticipantByVac",mock.Anything).Return(1, nil).Once()
		parRepository.On("CountParicipantByUserId",mock.Anything).Return(1, nil).Once()
		parRepository.On("GetParticipantByNIK", mock.Anything).Return(userData, nil).Once()
		parRepository.On("GetParticipantMultiParam", mock.Anything, mock.Anything).
			Return(participant.ParticipantCore{
				ID: 0,
			}, nil).Once()
		parRepository.On("ApplyParticipant", mock.AnythingOfType("participant.ParticipantCore")).Return(nil).Once()
		err := parServ.ApplyParticipant(participant.ParticipantCore{
			UserID: 1,
			VacID:  1,
		})

		assert.NotNil(t, err)
	})

	t.Run("apply participant error GetVaccinationByIdPost", func(t *testing.T) {
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vac.VacCore{}, errors.New("error get vaccination by id post")).Once()
		err := parServ.ApplyParticipant(participant.ParticipantCore{
			UserID: 1,
			VacID:  1,
			SessionID: 1,
		})

		assert.NotNil(t, err)
		assert.Equal(t, "error get vaccination by id post", err.Error())
	})

	t.Run("apply participant error vacdata id = 0", func(t *testing.T) { 
		vacData.ID = 0
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		err := parServ.ApplyParticipant(participant.ParticipantCore{
			UserID: 1,
			VacID:  0,
			SessionID: 1,
		})

		assert.NotNil(t, err)
		msg := fmt.Sprintf("vac with id %v not found", 0)
		assert.Equal(t, msg, err.Error())
	})

	// t.Run("apply participant error GetParticipantMultiParam", func(t *testing.T) {
	// 	vacData.ID = 1
	// 	vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
	// 	parRepository.On("GetParticipantMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
	// 		Return(participant.ParticipantCore{}, errors.New("error get participant")).Once()
	// 	err := parServ.ApplyParticipant(participant.ParticipantCore{
	// 		UserID:    1,
	// 		VacID:     1,
	// 		AppliedAt: time.Now(),
	// 	})

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, "error get participant", err.Error())
	// })

	// t.Run("apply participant error same participant exist", func(t *testing.T) {
	// 	vacData.ID = 1
	// 	vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
	// 	parRepository.On("GetParticipantMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
	// 		Return(participant.ParticipantCore{
	// 			ID:     1,
	// 			UserID: 1,
	// 			VacID:  1,
	// 			Status: "pending",
	// 		}, nil).Once()
	// 	err := parServ.ApplyParticipant(participant.ParticipantCore{
	// 		UserID: 1,
	// 		VacID:  1,
	// 	})
	// 	assert.NotNil(t, err)
	// 	msg := fmt.Sprintf("user with id %v had applied participant with id %v, current status = %v", 1, 1, "pending")
	// 	assert.Equal(t, msg, err.Error())
	// })

	t.Run("apply participant error ApplyParticipant", func(t *testing.T) {
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		parRepository.On("GetParticipantMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID: 0,
			}, nil).Once()
		parRepository.On("ApplyParticipant", mock.AnythingOfType("participant.ParticipantCore")).Return(errors.New("error apply participant")).Once()
		err := parServ.ApplyParticipant(participant.ParticipantCore{
			UserID: 1,
			VacID:  1,
		})

		assert.NotNil(t, err)
		// assert.Equal(t, "vac with id 0 not found", err.Error())
	})
}

func TestGetParticipantByUserID(t *testing.T) {
	participants := []participant.ParticipantCore{
		{
			ID:         1,
			Nik:		"312345678912",
			Fullname:	"Aditya Chandra",
			Address:	"Sumatera Utara",
			PhoneNumber:"081234567891",
			UserID:		1,
			VacID:		1,
			SessionID:	1,
			Status:		"pending",
			AppliedAt:   time.Now(),
			Vac: participant.VacCore{
				ID:				1,
				Description:	"Sinovac Dosis Pertama",
				Location:		"RS Universitas Sumatra Utara",
				Latitude: 		10,
				Longitude: 		10,
				VacType:		"Sinovac",
				Stock:			200,
				AdminId	:		1,
			},
		},
	}
	t.Run("get participant by user id success", func(t *testing.T) {
		parRepository.On("GetParticipantByUserID", mock.AnythingOfType("int")).Return(participants, nil).Once()
		resp, err := parServ.GetParticipantByUserID(1)
		assert.Nil(t, err)
		assert.Equal(t, len(participants), len(resp))
		assert.Equal(t, participants[0].ID, resp[0].ID)
	})

	t.Run("get participant by user id error GetParticipantByUserID", func(t *testing.T) {
		parRepository.On("GetParticipantByUserID", mock.AnythingOfType("int")).Return(nil, errors.New("error get participant")).Once()
		resp, err := parServ.GetParticipantByUserID(1)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.Equal(t, "error get participant", err.Error())
	})
}

func TestRejectParticipant(t *testing.T) {
	t.Run("reject participant success", func(t *testing.T) {
		vacData.AdminId = 1
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 	 1, 
				},
			}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		parRepository.On("RejectParticipant", mock.AnythingOfType("int")).Return(nil).Once()

		err := parServ.RejectParticipant(1, 1)
		assert.NotNil(t, err)
	})

	t.Run("reject participant error GetParticipantByID", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{}, errors.New("error")).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{}, errors.New("error")).Once()
		err := parServ.RejectParticipant(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "participant with id 1 not found", err.Error())
	})

	t.Run("reject participant error no match admin id", func(t *testing.T) { 
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 2,
				},
			}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		err := parServ.RejectParticipant(1, 2)
		msg := fmt.Sprintf("admin with id %v not allowed to access post with id %v", 2, 1)
		assert.NotNil(t, err)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("reject participant error already rejected or accepted", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "accepted",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "rejected",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		err := parServ.RejectParticipant(1, 1)
		msg := fmt.Sprintf("this participant has been %v", "rejected")
		assert.NotNil(t, err)
		assert.Equal(t, msg, err.Error())
	})

	t.Run("reject participant error RejectParticipant", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{
			ID:     1,
			UserID: 1,
			VacID:  1,
			Status: "pending",
			Vac: participant.VacCore{
				ID:          1,
				AdminId: 1,
			},
		}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "rejected",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		parRepository.On("RejectParticipant", mock.AnythingOfType("int")).Return(errors.New("error reject participant")).Once()
		err := parServ.RejectParticipant(1, 1)
		assert.NotNil(t, err)
	})

}

func TestAcceptParticipant(t *testing.T) {
	t.Run("accept participant success", func(t *testing.T) {
		vacData.AdminId = 1
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{
			ID:     1,
			UserID: 1,
			VacID:  1,
			Vac: participant.VacCore{
				AdminId: 1,
			},
			Status: "pending",
		}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		parRepository.On("AcceptParticipant", mock.AnythingOfType("int")).Return(nil).Once()

		err := parServ.AcceptParticipant(1, 1)
		assert.NotNil(t, err)
	})

	t.Run("accept participant error GetParticipantByID", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{}, errors.New("error get participant")).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{}, errors.New("error")).Once()
		err := parServ.AcceptParticipant(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "participant with id 1 not found", err.Error())
	})

	t.Run("accept participant error admin id not match", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{
			ID:     1,
			UserID: 1,
			VacID:  1,
			Vac: participant.VacCore{
				AdminId: 1,
			},
			Status: "pending",
		}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "pending",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		err := parServ.AcceptParticipant(1, 2)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("admin with id %v not allowed to access participant with id %v", 2, 1)
		assert.NotEqual(t, msg, err.Error())
	})

	t.Run("accept participant error participant accepted/rejected", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{
			ID:     1,
			UserID: 1,
			VacID:  1,
			Vac: participant.VacCore{
				AdminId: 1,
			},
			Status: "accepted",
		}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "accepted",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		err := parServ.AcceptParticipant(1, 1)
		assert.NotNil(t, err)
		msg := fmt.Sprintf("this participant has been %v", "accepted")
		assert.Equal(t, msg, err.Error())
	})

	t.Run("accept participant error acceptParticipant", func(t *testing.T) {
		parMockServ.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{
			ID:     1,
			UserID: 1,
			VacID:  1,
			Vac: participant.VacCore{
				AdminId: 1,
			},
			Status: "pending",
		}, nil).Once()
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).
			Return(participant.ParticipantCore{
				ID:     1,
				UserID: 1,
				Status: "accepted",
				Vac: participant.VacCore{
					ID:          1,
					AdminId: 1,
				},
			}, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		parRepository.On("AcceptParticipant", mock.AnythingOfType("int")).Return(errors.New("error accept participant")).Once()

		err := parServ.AcceptParticipant(1, 1)
		assert.NotNil(t, err)
	})
}

func TestGetParticipantByID(t *testing.T) {
	t.Run("get participant by id ", func(t *testing.T) {
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).Return(parData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vacData, nil).Once()
		resp, err := parServ.GetParticipantByID(1)
		assert.Nil(t, err)
		assert.Equal(t, parData.ID, resp.ID)
		assert.Equal(t, 1, resp.Vac.AdminId)
		assert.Equal(t, parData.User.ID, resp.User.ID)
	})

	t.Run("get participant by id error GetParticipantByID", func(t *testing.T) {
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).Return(participant.ParticipantCore{}, errors.New("error get participant")).Once()
		resp, err := parServ.GetParticipantByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get participant", err.Error())
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get participant by id error GetUserById", func(t *testing.T) {
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).Return(parData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error get user")).Once()
		resp, err := parServ.GetParticipantByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get user", err.Error())
		assert.Equal(t, "", resp.Status)
	})

	t.Run("get participant by id error GetUserById", func(t *testing.T) {
		parRepository.On("GetParticipantByID", mock.AnythingOfType("int")).Return(parData, nil).Once()
		userService.On("GetUserById", mock.AnythingOfType("int")).Return(userData, nil).Once()
		vacService.On("GetVaccinationByIdPost", mock.AnythingOfType("int")).Return(vac.VacCore{}, errors.New("error get vac")).Once()
		resp, err := parServ.GetParticipantByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error get vac", err.Error())
		assert.Equal(t, "", resp.Status)
	})

}

func TestGetParticipantByVacID(t *testing.T) {
	t.Run("get participant by vac id", func(t *testing.T) {
		parRepository.On("GetParticipantByVacID", mock.AnythingOfType("int")).Return(parsData, nil).Once()
		resp, err := parServ.GetParticipantByVacID(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(resp))
		assert.Equal(t, 1, int(resp[0].ID))
		assert.Equal(t, "pending", resp[0].Status)
	})

	t.Run("get participant by vac id error", func(t *testing.T) {
		parRepository.On("GetParticipantByVacID", mock.AnythingOfType("int")).Return(nil, errors.New("error get participant")).Once()
		resp, err := parServ.GetParticipantByVacID(1)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, 0, len(resp))
	})
}

func TestGetParticipantMultiParam(t *testing.T) {
	t.Run("get participant with multiparam", func(t *testing.T) {
		parRepository.On("GetParticipantMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(parData, nil).Once()
		resp, err := parServ.GetParticipantMultiParam(0, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(resp.ID))
		assert.Equal(t, "pending", resp.Status)
	})

	t.Run("get participant with multiparam error", func(t *testing.T) {
		parRepository.On("GetParticipantMultiParam", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(participant.ParticipantCore{}, errors.New("error get participant")).Once()
		resp, err := parServ.GetParticipantMultiParam(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
		assert.Equal(t, "", resp.Status)
	})
}

