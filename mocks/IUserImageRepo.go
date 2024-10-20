// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "project-skbackend/internal/models"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IUserImageRepo is an autogenerated mock type for the IUserImageRepo type
type IUserImageRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ui
func (_m *IUserImageRepo) Create(ui models.UserImage) (*models.UserImage, error) {
	ret := _m.Called(ui)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.UserImage
	var r1 error
	if rf, ok := ret.Get(0).(func(models.UserImage) (*models.UserImage, error)); ok {
		return rf(ui)
	}
	if rf, ok := ret.Get(0).(func(models.UserImage) *models.UserImage); ok {
		r0 = rf(ui)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserImage)
		}
	}

	if rf, ok := ret.Get(1).(func(models.UserImage) error); ok {
		r1 = rf(ui)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ui
func (_m *IUserImageRepo) Delete(ui models.UserImage) error {
	ret := _m.Called(ui)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.UserImage) error); ok {
		r0 = rf(ui)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *IUserImageRepo) GetByID(id uuid.UUID) (*models.UserImage, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *models.UserImage
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.UserImage, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.UserImage); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserImage)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: uid
func (_m *IUserImageRepo) GetByUserID(uid uuid.UUID) (*models.UserImage, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserID")
	}

	var r0 *models.UserImage
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.UserImage, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.UserImage); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserImage)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields:
func (_m *IUserImageRepo) Read() ([]*models.UserImage, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*models.UserImage
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.UserImage, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.UserImage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserImage)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ui
func (_m *IUserImageRepo) Update(ui models.UserImage) (*models.UserImage, error) {
	ret := _m.Called(ui)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.UserImage
	var r1 error
	if rf, ok := ret.Get(0).(func(models.UserImage) (*models.UserImage, error)); ok {
		return rf(ui)
	}
	if rf, ok := ret.Get(0).(func(models.UserImage) *models.UserImage); ok {
		r0 = rf(ui)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserImage)
		}
	}

	if rf, ok := ret.Get(1).(func(models.UserImage) error); ok {
		r1 = rf(ui)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserImageRepo creates a new instance of IUserImageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserImageRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserImageRepo {
	mock := &IUserImageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
