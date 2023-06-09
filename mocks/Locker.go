// Code generated by mockery v2.25.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Locker is an autogenerated mock type for the Locker type
type Locker struct {
	mock.Mock
}

type Locker_Expecter struct {
	mock *mock.Mock
}

func (_m *Locker) EXPECT() *Locker_Expecter {
	return &Locker_Expecter{mock: &_m.Mock}
}

// Lock provides a mock function with given fields:
func (_m *Locker) Lock() {
	_m.Called()
}

// Locker_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type Locker_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
func (_e *Locker_Expecter) Lock() *Locker_Lock_Call {
	return &Locker_Lock_Call{Call: _e.mock.On("Lock")}
}

func (_c *Locker_Lock_Call) Run(run func()) *Locker_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Locker_Lock_Call) Return() *Locker_Lock_Call {
	_c.Call.Return()
	return _c
}

func (_c *Locker_Lock_Call) RunAndReturn(run func()) *Locker_Lock_Call {
	_c.Call.Return(run)
	return _c
}

// Unlock provides a mock function with given fields:
func (_m *Locker) Unlock() {
	_m.Called()
}

// Locker_Unlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlock'
type Locker_Unlock_Call struct {
	*mock.Call
}

// Unlock is a helper method to define mock.On call
func (_e *Locker_Expecter) Unlock() *Locker_Unlock_Call {
	return &Locker_Unlock_Call{Call: _e.mock.On("Unlock")}
}

func (_c *Locker_Unlock_Call) Run(run func()) *Locker_Unlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Locker_Unlock_Call) Return() *Locker_Unlock_Call {
	_c.Call.Return()
	return _c
}

func (_c *Locker_Unlock_Call) RunAndReturn(run func()) *Locker_Unlock_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLocker interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocker creates a new instance of Locker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocker(t mockConstructorTestingTNewLocker) *Locker {
	mock := &Locker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
