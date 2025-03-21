// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ITelegramService is an autogenerated mock type for the ITelegramService type
type ITelegramService struct {
	mock.Mock
}

// SendMessage provides a mock function with given fields: msg
func (_m *ITelegramService) SendMessage(msg string) error {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewITelegramService creates a new instance of ITelegramService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITelegramService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITelegramService {
	mock := &ITelegramService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
