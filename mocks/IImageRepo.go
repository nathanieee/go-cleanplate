// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "project-skbackend/internal/models"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IImageRepo is an autogenerated mock type for the IImageRepo type
type IImageRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: i
func (_m *IImageRepo) Create(i models.Image) (*models.Image, error) {
	ret := _m.Called(i)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Image) (*models.Image, error)); ok {
		return rf(i)
	}
	if rf, ok := ret.Get(0).(func(models.Image) *models.Image); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Image) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: i
func (_m *IImageRepo) Delete(i models.Image) error {
	ret := _m.Called(i)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Image) error); ok {
		r0 = rf(i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *IImageRepo) GetByID(id uuid.UUID) (*models.Image, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *models.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.Image, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Image); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields:
func (_m *IImageRepo) Read() ([]*models.Image, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*models.Image
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Image, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Image)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: i
func (_m *IImageRepo) Update(i models.Image) (*models.Image, error) {
	ret := _m.Called(i)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Image) (*models.Image, error)); ok {
		return rf(i)
	}
	if rf, ok := ret.Get(0).(func(models.Image) *models.Image); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Image) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIImageRepo creates a new instance of IImageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIImageRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IImageRepo {
	mock := &IImageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
