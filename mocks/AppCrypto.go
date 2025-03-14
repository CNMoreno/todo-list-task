// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AppCrypto is an autogenerated mock type for the AppCrypto type
type AppCrypto struct {
	mock.Mock
}

// CompareHashAndPassword provides a mock function with given fields: hashedPassword, password
func (_m *AppCrypto) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	ret := _m.Called(hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for CompareHashAndPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateFromPassword provides a mock function with given fields: password, cost
func (_m *AppCrypto) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	ret := _m.Called(password, cost)

	if len(ret) == 0 {
		panic("no return value specified for GenerateFromPassword")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, int) ([]byte, error)); ok {
		return rf(password, cost)
	}
	if rf, ok := ret.Get(0).(func([]byte, int) []byte); ok {
		r0 = rf(password, cost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, int) error); ok {
		r1 = rf(password, cost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAppCrypto creates a new instance of AppCrypto. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppCrypto(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppCrypto {
	mock := &AppCrypto{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
