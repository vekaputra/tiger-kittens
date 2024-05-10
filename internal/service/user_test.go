package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/jwt"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	mockuser "github.com/vekaputra/tiger-kittens/internal/repository/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Register(t *testing.T) {
	type TestSuite struct {
		ctx                context.Context
		mockUserRepository *mockuser.UserRepositoryProvider
		payload            model.RegisterUserRequest
		service            *UserService
	}

	setupTest := func() *TestSuite {
		timeNow, _ := time.Parse(time.RFC3339, "2024-01-01T12:33:12Z00:00")
		payload := model.RegisterUserRequest{
			Email:    "test@mail.com",
			Password: "password",
			Username: "test-username",
		}

		mockUserRepository := &mockuser.UserRepositoryProvider{}
		service := NewUserService(UserConfig{}, mockUserRepository)
		service.fnTimeNow = func() time.Time {
			return timeNow
		}

		return &TestSuite{
			ctx:                context.Background(),
			mockUserRepository: mockUserRepository,
			payload:            payload,
			service:            service,
		}
	}

	t.Run("success", func(t *testing.T) {
		suite := setupTest()
		expectedResult := model.MessageResponse{
			Message:   _const.RegisterSuccessMessage,
			Timestamp: suite.service.fnTimeNow().Format(time.RFC3339),
		}

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Email, suite.payload.Username).Return([]entity.User{}, nil)
		suite.mockUserRepository.On("Insert", suite.ctx, mock.AnythingOfType("entity.User")).Return(nil)

		resp, err := suite.service.Register(suite.ctx, suite.payload)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on FindByEmailOrUsername failed", func(t *testing.T) {
		suite := setupTest()
		expectedErr := customerror.ErrorInternalServer

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Email, suite.payload.Username).Return([]entity.User{}, expectedErr)
		suite.mockUserRepository.AssertNotCalled(t, "Insert", suite.ctx, mock.AnythingOfType("entity.User"))

		resp, err := suite.service.Register(suite.ctx, suite.payload)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.MessageResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on FindByEmailOrUsername return non zero length result", func(t *testing.T) {
		suite := setupTest()
		respUsers := []entity.User{
			{
				ID:        "32778434-4c34-4fd8-8108-07f38719f633",
				Email:     "test@mail.com",
				Password:  "StrongPassword",
				Username:  "other-username",
				CreatedAt: time.Now(),
			},
		}

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Email, suite.payload.Username).Return(respUsers, nil)
		suite.mockUserRepository.AssertNotCalled(t, "Insert", suite.ctx, mock.AnythingOfType("entity.User"))

		resp, err := suite.service.Register(suite.ctx, suite.payload)

		assert.EqualError(t, customerror.ErrorDuplicateUser, err.Error())
		assert.Equal(t, model.MessageResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on password too long (more than 72 char)", func(t *testing.T) {
		suite := setupTest()
		invalidPasswordPayload := suite.payload
		invalidPasswordPayload.Password = "12345678901234567890123456789012345678901234567890123456789012345678901234567890"

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Email, suite.payload.Username).Return([]entity.User{}, nil)
		suite.mockUserRepository.AssertNotCalled(t, "Insert", suite.ctx, mock.AnythingOfType("entity.User"))

		resp, err := suite.service.Register(suite.ctx, invalidPasswordPayload)

		assert.EqualError(t, bcrypt.ErrPasswordTooLong, err.Error())
		assert.Equal(t, model.MessageResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on Insert failed", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Email, suite.payload.Username).Return([]entity.User{}, nil)
		suite.mockUserRepository.On("Insert", suite.ctx, mock.AnythingOfType("entity.User")).Return(customerror.ErrorInternalServer)

		resp, err := suite.service.Register(suite.ctx, suite.payload)

		assert.EqualError(t, customerror.ErrorInternalServer, err.Error())
		assert.Equal(t, model.MessageResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})
}

func TestUserService_Login(t *testing.T) {
	type TestSuite struct {
		ctx                context.Context
		mockUserRepository *mockuser.UserRepositoryProvider
		payload            model.LoginUserRequest
		respUsers          []entity.User
		service            *UserService
	}

	setupTest := func() *TestSuite {
		timeNow, _ := time.Parse(time.RFC3339, "2024-01-01T12:33:12Z00:00")
		payload := model.LoginUserRequest{
			Password: "password",
			Username: "test-username",
		}

		privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
		config := UserConfig{
			PrivateKey:           privateKey,
			ExpiredAfterInSecond: time.Hour,
		}

		mockUserRepository := &mockuser.UserRepositoryProvider{}
		service := NewUserService(config, mockUserRepository)
		service.fnTimeNow = func() time.Time {
			return timeNow
		}

		return &TestSuite{
			ctx:                context.Background(),
			mockUserRepository: mockUserRepository,
			payload:            payload,
			respUsers: []entity.User{
				{
					ID:        "32778434-4c34-4fd8-8108-07f38719f633",
					Email:     "test@mail.com",
					Password:  "$2a$12$3Zylff8cORP2x/zrXQsJa.EuNASXAK/yVmP21UDCPeT0fUNXICX02",
					Username:  "test-username",
					CreatedAt: time.Now(),
				},
			},
			service: service,
		}
	}

	t.Run("success", func(t *testing.T) {
		suite := setupTest()
		expectedJWT, _ := jwt.GenerateAccessToken(suite.service.Config.PrivateKey, suite.service.fnTimeNow().Add(suite.service.Config.ExpiredAfterInSecond).Unix(), suite.respUsers[0])
		expectedResult := model.LoginUserResponse{
			AccessToken: expectedJWT,
			Timestamp:   suite.service.fnTimeNow().Format(time.RFC3339),
		}

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Username, suite.payload.Username).Return(suite.respUsers, nil)

		resp, err := suite.service.Login(suite.ctx, suite.payload)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on FindByEmailOrUsername failed", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Username, suite.payload.Username).Return([]entity.User{}, customerror.ErrorInternalServer)

		resp, err := suite.service.Login(suite.ctx, suite.payload)

		assert.EqualError(t, customerror.ErrorInternalServer, err.Error())
		assert.Equal(t, model.LoginUserResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on FindByEmailOrUsername return zero result", func(t *testing.T) {
		suite := setupTest()

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Username, suite.payload.Username).Return([]entity.User{}, nil)

		resp, err := suite.service.Login(suite.ctx, suite.payload)

		assert.EqualError(t, customerror.ErrorInvalidCredential, err.Error())
		assert.Equal(t, model.LoginUserResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})

	t.Run("failed on invalid password", func(t *testing.T) {
		suite := setupTest()
		invalidPasswordPayload := suite.payload
		invalidPasswordPayload.Password = "invalid-password"

		suite.mockUserRepository.On("FindByEmailOrUsername", suite.ctx, suite.payload.Username, suite.payload.Username).Return(suite.respUsers, nil)

		resp, err := suite.service.Login(suite.ctx, invalidPasswordPayload)

		assert.EqualError(t, customerror.ErrorInvalidCredential, err.Error())
		assert.Equal(t, model.LoginUserResponse{}, resp)
		suite.mockUserRepository.AssertExpectations(t)
	})
}
