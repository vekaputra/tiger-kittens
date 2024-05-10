package pgsql

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	sqlxmock "github.com/vekaputra/tiger-kittens/pkg/database/sqlx/mock"
)

func TestUserRepository_FindByEmailOrUsername(t *testing.T) {
	ctx := context.Background()

	dbMock := dbTest
	_, err := dbMock.Exec(`
INSERT INTO users (id, username, email, password) VALUES 
	('32778434-4c34-4fd8-8108-07f38719f633', 'test-username-1', 'test1@mail.com', '$2a$12$MGpAKLP3TiaALlGuLUTs7eIBXLt4aVaZCm.DXKySziPlOMcCNU5Va'),
	('e5663f77-0e5e-4b47-badd-215d32bdf195', 'test-username-2', 'test2@mail.com', '$2a$12$MGpAKLP3TiaALlGuLUTs7eIBXLt4aVaZCm.DXKySziPlOMcCNU5Va'),
	('1a41aeb5-513c-4c79-9d37-812760e17c8b', 'test-username-3', 'test3@mail.com', '$2a$12$MGpAKLP3TiaALlGuLUTs7eIBXLt4aVaZCm.DXKySziPlOMcCNU5Va');
`)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	defer func() {
		_, err = dbMock.Exec(`TRUNCATE users CASCADE;`)
		if err != nil {
			assert.Fail(t, err.Error())
		}
	}()

	repo := NewUserRepository(dbMock)

	t.Run("success find email", func(t *testing.T) {
		result, err := repo.FindByEmailOrUsername(ctx, "test1@mail.com", "test1@mail.com")

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "32778434-4c34-4fd8-8108-07f38719f633", result[0].ID)
	})

	t.Run("success find username", func(t *testing.T) {
		result, err := repo.FindByEmailOrUsername(ctx, "test-username-2", "test-username-2")

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "e5663f77-0e5e-4b47-badd-215d32bdf195", result[0].ID)
	})

	t.Run("success find with both email and username", func(t *testing.T) {
		result, err := repo.FindByEmailOrUsername(ctx, "test1@mail.com", "test-username-2")

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "32778434-4c34-4fd8-8108-07f38719f633", result[0].ID)
		assert.Equal(t, "e5663f77-0e5e-4b47-badd-215d32bdf195", result[1].ID)
	})

	t.Run("success with invalid email", func(t *testing.T) {
		result, err := repo.FindByEmailOrUsername(ctx, "not-found@mail.com", "not-found@mail.com")

		assert.NoError(t, err)
		assert.Len(t, result, 0)
	})

	t.Run("success with invalid username", func(t *testing.T) {
		result, err := repo.FindByEmailOrUsername(ctx, "not-found", "not-found")

		assert.NoError(t, err)
		assert.Len(t, result, 0)
	})

	t.Run("failed if SelectContext failed", func(t *testing.T) {
		expectedErr := customerror.ErrorInternalServer

		failDBMock := &sqlxmock.DBer{}
		failDBMock.On("SelectContext", ctx, mock.Anything, mock.Anything, "test1@mail.com", "test1@mail.com").Return(expectedErr)

		failRepo := NewUserRepository(failDBMock)
		users, err := failRepo.FindByEmailOrUsername(ctx, "test1@mail.com", "test1@mail.com")

		assert.EqualError(t, expectedErr, err.Error())
		assert.Len(t, users, 0)
		failDBMock.AssertExpectations(t)
	})
}

func TestUserRepository_Insert(t *testing.T) {

	ctx := context.Background()

	dbMock := dbTest
	defer func() {
		_, err := dbMock.Exec(`TRUNCATE users CASCADE;`)
		if err != nil {
			assert.Fail(t, err.Error())
		}
	}()

	repo := NewUserRepository(dbMock)
	user := entity.User{
		Email:    "test@mail.com",
		Password: "password",
		Username: "test-username",
	}

	t.Run("success insert", func(t *testing.T) {
		err := repo.Insert(ctx, user)
		assert.NoError(t, err)

		var result entity.User
		err = dbMock.Get(&result, "SELECT id, email, password, username, created_at FROM users WHERE username = 'test-username';")
		assert.NoError(t, err)

		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.Password, result.Password)
		assert.Equal(t, user.Username, result.Username)
	})

	t.Run("failed if ExecContext failed", func(t *testing.T) {
		expectedErr := customerror.ErrorInternalServer

		failDBMock := &sqlxmock.DBer{}
		failDBMock.On("ExecContext", ctx, mock.Anything, mock.Anything, user.Email, user.Password, user.Username).Return(driver.RowsAffected(0), expectedErr)

		failRepo := NewUserRepository(failDBMock)
		err := failRepo.Insert(ctx, user)

		assert.EqualError(t, expectedErr, err.Error())
		failDBMock.AssertExpectations(t)
	})
}
