// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AuthStore is an autogenerated mock type for the AuthStore type
type AuthStore struct {
	mock.Mock
}

// ContextForNamespace provides a mock function with given fields: baseCtx, namespace
func (_m *AuthStore) ContextForNamespace(baseCtx context.Context, namespace string) (context.Context, error) {
	ret := _m.Called(baseCtx, namespace)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(baseCtx, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(baseCtx, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveAllCreatedTokens provides a mock function with given fields: ctx
func (_m *AuthStore) RemoveAllCreatedTokens(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}