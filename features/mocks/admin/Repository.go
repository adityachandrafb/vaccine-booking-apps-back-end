// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	admin "vac/features/admin"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CheckAdmin provides a mock function with given fields: data
func (_m *Repository) CheckAdmin(data admin.AdminCore) (admin.AdminCore, error) {
	ret := _m.Called(data)

	var r0 admin.AdminCore
	if rf, ok := ret.Get(0).(func(admin.AdminCore) admin.AdminCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(admin.AdminCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(admin.AdminCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAdmin provides a mock function with given fields: data
func (_m *Repository) CreateAdmin(data admin.AdminCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(admin.AdminCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdminById provides a mock function with given fields: data
func (_m *Repository) GetAdminById(data admin.AdminCore) (admin.AdminCore, error) {
	ret := _m.Called(data)

	var r0 admin.AdminCore
	if rf, ok := ret.Get(0).(func(admin.AdminCore) admin.AdminCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(admin.AdminCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(admin.AdminCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdmins provides a mock function with given fields:
func (_m *Repository) GetAdmins() ([]admin.AdminCore, error) {
	ret := _m.Called()

	var r0 []admin.AdminCore
	if rf, ok := ret.Get(0).(func() []admin.AdminCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admin.AdminCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
