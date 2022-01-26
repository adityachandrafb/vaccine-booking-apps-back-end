// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	vac "vac/features/vac"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteVacData provides a mock function with given fields: data
func (_m *Repository) DeleteVacData(data vac.VacCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(vac.VacCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNearbyFacilities provides a mock function with given fields: latitude, longitude, radius
func (_m *Repository) GetNearbyFacilities(latitude float64, longitude float64, radius float64) ([]vac.VacCore, error) {
	ret := _m.Called(latitude, longitude, radius)

	var r0 []vac.VacCore
	if rf, ok := ret.Get(0).(func(float64, float64, float64) []vac.VacCore); ok {
		r0 = rf(latitude, longitude, radius)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vac.VacCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float64, float64, float64) error); ok {
		r1 = rf(latitude, longitude, radius)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVacByIdAdmin provides a mock function with given fields: id
func (_m *Repository) GetVacByIdAdmin(id int) ([]vac.VacCore, error) {
	ret := _m.Called(id)

	var r0 []vac.VacCore
	if rf, ok := ret.Get(0).(func(int) []vac.VacCore); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vac.VacCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVacData provides a mock function with given fields: data
func (_m *Repository) GetVacData(data vac.VacCore) ([]vac.VacCore, error) {
	ret := _m.Called(data)

	var r0 []vac.VacCore
	if rf, ok := ret.Get(0).(func(vac.VacCore) []vac.VacCore); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vac.VacCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(vac.VacCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVacDataById provides a mock function with given fields: id
func (_m *Repository) GetVacDataById(id int) (vac.VacCore, error) {
	ret := _m.Called(id)

	var r0 vac.VacCore
	if rf, ok := ret.Get(0).(func(int) vac.VacCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vac.VacCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertData provides a mock function with given fields: data
func (_m *Repository) InsertData(data vac.VacCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(vac.VacCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateVacData provides a mock function with given fields: data
func (_m *Repository) UpdateVacData(data vac.VacCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(vac.VacCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
