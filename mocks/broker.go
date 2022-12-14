// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	chat "github.com/guil95/chat-go/internal/chat"
	mock "github.com/stretchr/testify/mock"
)

// Broker is an autogenerated mock type for the Broker type
type Broker struct {
	mock.Mock
}

// Consume provides a mock function with given fields: messageReceived
func (_m *Broker) Consume(messageReceived chan []byte) error {
	ret := _m.Called(messageReceived)

	var r0 error
	if rf, ok := ret.Get(0).(func(chan []byte) error); ok {
		r0 = rf(messageReceived)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *Broker) Send(_a0 *chat.Chat) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*chat.Chat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBroker interface {
	mock.TestingT
	Cleanup(func())
}

// NewBroker creates a new instance of Broker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBroker(t mockConstructorTestingTNewBroker) *Broker {
	mock := &Broker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
