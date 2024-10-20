// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "project-skbackend/internal/models"

	mock "github.com/stretchr/testify/mock"

	utpagination "project-skbackend/packages/utils/utpagination"

	uuid "github.com/google/uuid"
)

// IAdminRepository is an autogenerated mock type for the IAdminRepository type
type IAdminRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: a
func (_m *IAdminRepository) Create(a models.Admin) (*models.Admin, error) {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Admin) (*models.Admin, error)); ok {
		return rf(a)
	}
	if rf, ok := ret.Get(0).(func(models.Admin) *models.Admin); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Admin)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Admin) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: a
func (_m *IAdminRepository) Delete(a models.Admin) error {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Admin) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: p
func (_m *IAdminRepository) FindAll(p utpagination.Pagination) (*utpagination.Pagination, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 *utpagination.Pagination
	var r1 error
	if rf, ok := ret.Get(0).(func(utpagination.Pagination) (*utpagination.Pagination, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(utpagination.Pagination) *utpagination.Pagination); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*utpagination.Pagination)
		}
	}

	if rf, ok := ret.Get(1).(func(utpagination.Pagination) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *IAdminRepository) GetByEmail(email string) (*models.Admin, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *models.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Admin, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Admin); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Admin)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *IAdminRepository) GetByID(id uuid.UUID) (*models.Admin, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *models.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.Admin, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Admin); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Admin)
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
func (_m *IAdminRepository) GetByUserID(uid uuid.UUID) (*models.Admin, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserID")
	}

	var r0 *models.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.Admin, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Admin); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Admin)
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
func (_m *IAdminRepository) Read() ([]*models.Admin, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*models.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*models.Admin, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*models.Admin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Admin)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: a
func (_m *IAdminRepository) Update(a models.Admin) (*models.Admin, error) {
	ret := _m.Called(a)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.Admin
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Admin) (*models.Admin, error)); ok {
		return rf(a)
	}
	if rf, ok := ret.Get(0).(func(models.Admin) *models.Admin); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Admin)
		}
	}

	if rf, ok := ret.Get(1).(func(models.Admin) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIAdminRepository creates a new instance of IAdminRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAdminRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAdminRepository {
	mock := &IAdminRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
