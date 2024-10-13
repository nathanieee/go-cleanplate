// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	requests "project-skbackend/internal/controllers/requests"
	responses "project-skbackend/internal/controllers/responses"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ICartService is an autogenerated mock type for the ICartService type
type ICartService struct {
	mock.Mock
}

// Create provides a mock function with given fields: req, roleres
func (_m *ICartService) Create(req requests.CreateCart, roleres responses.BaseRole) (*responses.Cart, error) {
	ret := _m.Called(req, roleres)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *responses.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(requests.CreateCart, responses.BaseRole) (*responses.Cart, error)); ok {
		return rf(req, roleres)
	}
	if rf, ok := ret.Get(0).(func(requests.CreateCart, responses.BaseRole) *responses.Cart); ok {
		r0 = rf(req, roleres)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(requests.CreateCart, responses.BaseRole) error); ok {
		r1 = rf(req, roleres)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ICartService) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *ICartService) GetByID(id uuid.UUID) (*responses.Cart, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *responses.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*responses.Cart, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *responses.Cart); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.Cart)
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
func (_m *ICartService) Read() ([]*responses.Cart, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*responses.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*responses.Cart, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*responses.Cart); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*responses.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: cid, req
func (_m *ICartService) Update(cid uuid.UUID, req requests.UpdateCart) (*responses.Cart, error) {
	ret := _m.Called(cid, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *responses.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, requests.UpdateCart) (*responses.Cart, error)); ok {
		return rf(cid, req)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, requests.UpdateCart) *responses.Cart); ok {
		r0 = rf(cid, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, requests.UpdateCart) error); ok {
		r1 = rf(cid, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICartService creates a new instance of ICartService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICartService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICartService {
	mock := &ICartService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
