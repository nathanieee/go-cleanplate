// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	requests "project-skbackend/internal/controllers/requests"
	responses "project-skbackend/internal/controllers/responses"

	mock "github.com/stretchr/testify/mock"

	utpagination "project-skbackend/packages/utils/utpagination"

	uuid "github.com/google/uuid"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: req
func (_m *IUserService) Create(req requests.CreateUser) (*responses.User, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *responses.User
	var r1 error
	if rf, ok := ret.Get(0).(func(requests.CreateUser) (*responses.User, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(requests.CreateUser) *responses.User); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.User)
		}
	}

	if rf, ok := ret.Get(1).(func(requests.CreateUser) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: uid
func (_m *IUserService) Delete(uid uuid.UUID) error {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: p
func (_m *IUserService) FindAll(p utpagination.Pagination) (*utpagination.Pagination, error) {
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

// GetByID provides a mock function with given fields: uid
func (_m *IUserService) GetByID(uid uuid.UUID) (*responses.User, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *responses.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*responses.User, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *responses.User); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleDataByUserID provides a mock function with given fields: uid
func (_m *IUserService) GetRoleDataByUserID(uid uuid.UUID) (*responses.BaseRole, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetRoleDataByUserID")
	}

	var r0 *responses.BaseRole
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*responses.BaseRole, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *responses.BaseRole); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.BaseRole)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserName provides a mock function with given fields: uid
func (_m *IUserService) GetUserName(uid uuid.UUID) (string, string, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetUserName")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (string, string, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) string); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) string); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(uuid.UUID) error); ok {
		r2 = rf(uid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: req, uid
func (_m *IUserService) Update(req requests.UpdateUser, uid uuid.UUID) (*responses.User, error) {
	ret := _m.Called(req, uid)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *responses.User
	var r1 error
	if rf, ok := ret.Get(0).(func(requests.UpdateUser, uuid.UUID) (*responses.User, error)); ok {
		return rf(req, uid)
	}
	if rf, ok := ret.Get(0).(func(requests.UpdateUser, uuid.UUID) *responses.User); ok {
		r0 = rf(req, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.User)
		}
	}

	if rf, ok := ret.Get(1).(func(requests.UpdateUser, uuid.UUID) error); ok {
		r1 = rf(req, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOwnPassword provides a mock function with given fields: uid, req
func (_m *IUserService) UpdateOwnPassword(uid uuid.UUID, req requests.UpdatePassword) error {
	ret := _m.Called(uid, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOwnPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, requests.UpdatePassword) error); ok {
		r0 = rf(uid, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUserService creates a new instance of IUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserService {
	mock := &IUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
