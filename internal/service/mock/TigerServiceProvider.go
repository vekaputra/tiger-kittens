// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/vekaputra/tiger-kittens/internal/model"
)

// TigerServiceProvider is an autogenerated mock type for the TigerServiceProvider type
type TigerServiceProvider struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, payload
func (_m *TigerServiceProvider) Create(ctx context.Context, payload model.CreateTigerRequest) (model.MessageResponse, error) {
	ret := _m.Called(ctx, payload)

	var r0 model.MessageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateTigerRequest) (model.MessageResponse, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateTigerRequest) model.MessageResponse); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(model.MessageResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateTigerRequest) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSighting provides a mock function with given fields: ctx, payload
func (_m *TigerServiceProvider) CreateSighting(ctx context.Context, payload model.CreateSightingRequest) (model.MessageResponse, error) {
	ret := _m.Called(ctx, payload)

	var r0 model.MessageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateSightingRequest) (model.MessageResponse, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateSightingRequest) model.MessageResponse); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(model.MessageResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateSightingRequest) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, page
func (_m *TigerServiceProvider) List(ctx context.Context, page model.PaginationRequest) (model.ListTigerResponse, error) {
	ret := _m.Called(ctx, page)

	var r0 model.ListTigerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) (model.ListTigerResponse, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) model.ListTigerResponse); ok {
		r0 = rf(ctx, page)
	} else {
		r0 = ret.Get(0).(model.ListTigerResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.PaginationRequest) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSighting provides a mock function with given fields: ctx, page
func (_m *TigerServiceProvider) ListSighting(ctx context.Context, page model.PaginationRequest) (model.ListSightingResponse, error) {
	ret := _m.Called(ctx, page)

	var r0 model.ListSightingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) (model.ListSightingResponse, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) model.ListSightingResponse); ok {
		r0 = rf(ctx, page)
	} else {
		r0 = ret.Get(0).(model.ListSightingResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.PaginationRequest) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTigerServiceProvider creates a new instance of TigerServiceProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTigerServiceProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TigerServiceProvider {
	mock := &TigerServiceProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
