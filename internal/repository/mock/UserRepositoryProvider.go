// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/vekaputra/tiger-kittens/internal/repository/entity"
)

// UserRepositoryProvider is an autogenerated mock type for the UserRepositoryProvider type
type UserRepositoryProvider struct {
	mock.Mock
}

// FindByEmailOrUsername provides a mock function with given fields: ctx, email, username
func (_m *UserRepositoryProvider) FindByEmailOrUsername(ctx context.Context, email string, username string) ([]entity.User, error) {
	ret := _m.Called(ctx, email, username)

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]entity.User, error)); ok {
		return rf(ctx, email, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []entity.User); ok {
		r0 = rf(ctx, email, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, _a1
func (_m *UserRepositoryProvider) Insert(ctx context.Context, _a1 entity.User) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepositoryProvider creates a new instance of UserRepositoryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryProvider {
	mock := &UserRepositoryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
