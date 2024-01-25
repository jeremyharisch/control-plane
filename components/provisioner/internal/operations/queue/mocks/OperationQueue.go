// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OperationQueue is an autogenerated mock type for the OperationQueue type
type OperationQueue struct {
	mock.Mock
}

// Add provides a mock function with given fields: processId
func (_m *OperationQueue) Add(processId string) {
	_m.Called(processId)
}

// Run provides a mock function with given fields: stop
func (_m *OperationQueue) Run(stop <-chan struct{}) {
	_m.Called(stop)
}

// NewOperationQueue creates a new instance of OperationQueue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOperationQueue(t interface {
	mock.TestingT
	Cleanup(func())
}) *OperationQueue {
	mock := &OperationQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}