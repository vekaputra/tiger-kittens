// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"
)

// TxProvider is an autogenerated mock type for the TxProvider type
type TxProvider struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx
func (_m *TxProvider) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 *sqlx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*sqlx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *sqlx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseTx provides a mock function with given fields: tx, err
func (_m *TxProvider) CloseTx(tx *sqlx.Tx, err error) error {
	ret := _m.Called(tx, err)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sqlx.Tx, error) error); ok {
		r0 = rf(tx, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTxProvider creates a new instance of TxProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxProvider {
	mock := &TxProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
