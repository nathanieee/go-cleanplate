// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	requests "project-skbackend/internal/controllers/requests"
)

// IProducerService is an autogenerated mock type for the IProducerService type
type IProducerService struct {
	mock.Mock
}

// PublishEmail provides a mock function with given fields: message
func (_m *IProducerService) PublishEmail(message requests.SendEmail) error {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for PublishEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(requests.SendEmail) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIProducerService creates a new instance of IProducerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProducerService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProducerService {
	mock := &IProducerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
