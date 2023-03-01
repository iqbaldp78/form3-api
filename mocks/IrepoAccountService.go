// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/iqbaldp78/form3/models"
	mock "github.com/stretchr/testify/mock"
)

// IrepoAccountService is an autogenerated mock type for the IrepoAccountService type
type IrepoAccountService struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: fullPath, payload
func (_m *IrepoAccountService) CreateAccount(fullPath string, payload models.AccountData) (models.AccountData, error) {
	ret := _m.Called(fullPath, payload)

	var r0 models.AccountData
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.AccountData) (models.AccountData, error)); ok {
		return rf(fullPath, payload)
	}
	if rf, ok := ret.Get(0).(func(string, models.AccountData) models.AccountData); ok {
		r0 = rf(fullPath, payload)
	} else {
		r0 = ret.Get(0).(models.AccountData)
	}

	if rf, ok := ret.Get(1).(func(string, models.AccountData) error); ok {
		r1 = rf(fullPath, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccount provides a mock function with given fields: fullPath
func (_m *IrepoAccountService) DeleteAccount(fullPath string) error {
	ret := _m.Called(fullPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fullPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAccount provides a mock function with given fields: fullPath
func (_m *IrepoAccountService) FetchAccount(fullPath string) (models.AccountData, error) {
	ret := _m.Called(fullPath)

	var r0 models.AccountData
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.AccountData, error)); ok {
		return rf(fullPath)
	}
	if rf, ok := ret.Get(0).(func(string) models.AccountData); ok {
		r0 = rf(fullPath)
	} else {
		r0 = ret.Get(0).(models.AccountData)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fullPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIrepoAccountService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIrepoAccountService creates a new instance of IrepoAccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIrepoAccountService(t mockConstructorTestingTNewIrepoAccountService) *IrepoAccountService {
	mock := &IrepoAccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
