// Code generated by mockery v2.36.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/vekaputra/tiger-kittens/internal/repository/entity"

	model "github.com/vekaputra/tiger-kittens/internal/model"

	sqlx "github.com/jmoiron/sqlx"
)

// TigerRepositoryProvider is an autogenerated mock type for the TigerRepositoryProvider type
type TigerRepositoryProvider struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx
func (_m *TigerRepositoryProvider) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
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
func (_m *TigerRepositoryProvider) CloseTx(tx *sqlx.Tx, err error) error {
	ret := _m.Called(tx, err)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sqlx.Tx, error) error); ok {
		r0 = rf(tx, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: ctx
func (_m *TigerRepositoryProvider) Count(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountSighting provides a mock function with given fields: ctx
func (_m *TigerRepositoryProvider) CountSighting(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: ctx, name
func (_m *TigerRepositoryProvider) FindByName(ctx context.Context, name string) ([]entity.Tiger, error) {
	ret := _m.Called(ctx, name)

	var r0 []entity.Tiger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Tiger, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Tiger); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Tiger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSightingWithPagination provides a mock function with given fields: ctx, page
func (_m *TigerRepositoryProvider) FindSightingWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Sighting, error) {
	ret := _m.Called(ctx, page)

	var r0 []entity.Sighting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) ([]entity.Sighting, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) []entity.Sighting); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Sighting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.PaginationRequest) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWithPagination provides a mock function with given fields: ctx, page
func (_m *TigerRepositoryProvider) FindWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Tiger, error) {
	ret := _m.Called(ctx, page)

	var r0 []entity.Tiger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) ([]entity.Tiger, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.PaginationRequest) []entity.Tiger); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Tiger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.PaginationRequest) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, _a1
func (_m *TigerRepositoryProvider) Insert(ctx context.Context, _a1 entity.Tiger) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Tiger) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxFindByID provides a mock function with given fields: ctx, tx, id
func (_m *TigerRepositoryProvider) TxFindByID(ctx context.Context, tx *sqlx.Tx, id int) (*entity.Tiger, error) {
	ret := _m.Called(ctx, tx, id)

	var r0 *entity.Tiger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, int) (*entity.Tiger, error)); ok {
		return rf(ctx, tx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, int) *entity.Tiger); ok {
		r0 = rf(ctx, tx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Tiger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, int) error); ok {
		r1 = rf(ctx, tx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxInsertSighting provides a mock function with given fields: ctx, tx, _a2
func (_m *TigerRepositoryProvider) TxInsertSighting(ctx context.Context, tx *sqlx.Tx, _a2 entity.TigerSighting) error {
	ret := _m.Called(ctx, tx, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, entity.TigerSighting) error); ok {
		r0 = rf(ctx, tx, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxUpdate provides a mock function with given fields: ctx, tx, _a2
func (_m *TigerRepositoryProvider) TxUpdate(ctx context.Context, tx *sqlx.Tx, _a2 entity.Tiger) error {
	ret := _m.Called(ctx, tx, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, entity.Tiger) error); ok {
		r0 = rf(ctx, tx, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTigerRepositoryProvider creates a new instance of TigerRepositoryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTigerRepositoryProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TigerRepositoryProvider {
	mock := &TigerRepositoryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
