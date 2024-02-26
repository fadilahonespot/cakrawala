// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OpenAIWrapper is an autogenerated mock type for the OpenAIWrapper type
type OpenAIWrapper struct {
	mock.Mock
}

// GenerateText provides a mock function with given fields: ctx, sysText, userText
func (_m *OpenAIWrapper) GenerateText(ctx context.Context, sysText string, userText string) (string, error) {
	ret := _m.Called(ctx, sysText, userText)

	if len(ret) == 0 {
		panic("no return value specified for GenerateText")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, sysText, userText)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, sysText, userText)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sysText, userText)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOpenAIWrapper creates a new instance of OpenAIWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOpenAIWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *OpenAIWrapper {
	mock := &OpenAIWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
