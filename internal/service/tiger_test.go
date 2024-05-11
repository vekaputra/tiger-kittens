package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	mockrepo "github.com/vekaputra/tiger-kittens/internal/repository/mock"
)

func TestTigerService_List(t *testing.T) {
	type TestSuite struct {
		ctx                 context.Context
		mockTigerRepository *mockrepo.TigerRepositoryProvider
		service             *TigerService
	}

	setupTest := func() *TestSuite {
		timeNow, _ := time.Parse(time.RFC3339, "2024-01-01T12:33:12Z00:00")
		mockTigerRepository := &mockrepo.TigerRepositoryProvider{}
		service := NewTigerService(nil, mockTigerRepository)
		service.fnTimeNow = func() time.Time {
			return timeNow
		}

		return &TestSuite{
			ctx:                 context.Background(),
			mockTigerRepository: mockTigerRepository,
			service:             service,
		}
	}

	t.Run("success", func(t *testing.T) {
		suite := setupTest()
		page := model.PaginationRequest{
			Page:    1,
			PerPage: 3,
		}
		respCount := uint64(10)
		respTiger := []entity.Tiger{
			{
				ID:          1,
				DateOfBirth: "2024-01-01",
				LastLat:     50,
				LastLong:    50,
				LastSeen:    suite.service.fnTimeNow(),
				LastPhoto:   "public/no_image.svg",
				Name:        "tiger_1",
				CreatedAt:   suite.service.fnTimeNow(),
				UpdatedAt:   suite.service.fnTimeNow(),
			},
			{
				ID:          2,
				DateOfBirth: "2024-01-01",
				LastLat:     50,
				LastLong:    50,
				LastSeen:    suite.service.fnTimeNow(),
				LastPhoto:   "public/no_image.svg",
				Name:        "tiger_2",
				CreatedAt:   suite.service.fnTimeNow(),
				UpdatedAt:   suite.service.fnTimeNow(),
			},
			{
				ID:          3,
				DateOfBirth: "2024-01-01",
				LastLat:     50,
				LastLong:    50,
				LastSeen:    suite.service.fnTimeNow(),
				LastPhoto:   "public/no_image.svg",
				Name:        "tiger_3",
				CreatedAt:   suite.service.fnTimeNow(),
				UpdatedAt:   suite.service.fnTimeNow(),
			},
		}
		expectedResult := model.ListTigerResponse{
			Data: respTiger,
			Pagination: model.PaginationResponse{
				Page:      page.Page,
				PerPage:   page.PerPage,
				TotalPage: 4,
				TotalItem: respCount,
			},
		}

		suite.mockTigerRepository.On("Count", suite.ctx).Return(respCount, nil)
		suite.mockTigerRepository.On("FindWithPagination", suite.ctx, page).Return(respTiger, nil)

		result, err := suite.service.List(suite.ctx, page)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("success on page over total_page", func(t *testing.T) {
		suite := setupTest()
		page := model.PaginationRequest{
			Page:    5,
			PerPage: 3,
		}
		respCount := uint64(10)
		respTiger := []entity.Tiger{}
		expectedResult := model.ListTigerResponse{
			Data: respTiger,
			Pagination: model.PaginationResponse{
				Page:      page.Page,
				PerPage:   page.PerPage,
				TotalPage: 4,
				TotalItem: respCount,
			},
		}

		suite.mockTigerRepository.On("Count", suite.ctx).Return(respCount, nil)
		suite.mockTigerRepository.On("FindWithPagination", suite.ctx, page).Return(respTiger, nil)

		result, err := suite.service.List(suite.ctx, page)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on Count failed", func(t *testing.T) {
		suite := setupTest()
		page := model.PaginationRequest{
			Page:    1,
			PerPage: 3,
		}
		expectedErr := customerror.ErrorInternalServer

		suite.mockTigerRepository.On("Count", suite.ctx).Return(uint64(0), expectedErr)
		suite.mockTigerRepository.AssertNotCalled(t, "FindWithPagination", suite.ctx, page)

		result, err := suite.service.List(suite.ctx, page)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.ListTigerResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on FindWithPagination failed", func(t *testing.T) {
		suite := setupTest()
		page := model.PaginationRequest{
			Page:    1,
			PerPage: 3,
		}
		expectedErr := customerror.ErrorInternalServer

		suite.mockTigerRepository.On("Count", suite.ctx).Return(uint64(10), nil)
		suite.mockTigerRepository.On("FindWithPagination", suite.ctx, page).Return([]entity.Tiger{}, expectedErr)

		result, err := suite.service.List(suite.ctx, page)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.ListTigerResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})
}

func TestTigerService_Create(t *testing.T) {
	type TestSuite struct {
		ctx                 context.Context
		mockTigerRepository *mockrepo.TigerRepositoryProvider
		service             *TigerService
		payload             model.CreateTigerRequest
		payloadInsert       entity.Tiger
	}

	setupTest := func() *TestSuite {
		timeDoB, _ := time.Parse(time.DateOnly, "2024-01-01")
		timeNow, _ := time.Parse(time.RFC3339, "2024-01-01T12:33:12Z00:00")
		mockTigerRepository := &mockrepo.TigerRepositoryProvider{}
		service := NewTigerService(nil, mockTigerRepository)
		service.fnTimeNow = func() time.Time {
			return timeNow
		}

		return &TestSuite{
			ctx:                 context.Background(),
			mockTigerRepository: mockTigerRepository,
			service:             service,
			payload: model.CreateTigerRequest{
				LastLat:     50,
				LastLong:    50,
				LastSeen:    timeNow,
				Name:        "tiger_new",
				DateOfBirth: timeDoB,
				LastPhoto:   "public/no_image.svg",
			},
			payloadInsert: entity.Tiger{
				DateOfBirth: timeDoB.Format(time.DateOnly),
				LastLat:     50,
				LastLong:    50,
				LastSeen:    timeNow,
				LastPhoto:   "public/no_image.svg",
				Name:        "tiger_new",
			},
		}
	}

	t.Run("success", func(t *testing.T) {
		suite := setupTest()
		expectedResult := model.MessageResponse{
			Message:   _const.CreateTigerSuccessMessage,
			Timestamp: suite.service.fnTimeNow().Format(time.RFC3339),
		}

		suite.mockTigerRepository.On("FindByName", suite.ctx, suite.payload.Name).Return([]entity.Tiger{}, nil)
		suite.mockTigerRepository.On("Insert", suite.ctx, suite.payloadInsert).Return(nil)

		result, err := suite.service.Create(suite.ctx, suite.payload)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on FindByName failed", func(t *testing.T) {
		suite := setupTest()
		expectedErr := customerror.ErrorInternalServer

		suite.mockTigerRepository.On("FindByName", suite.ctx, suite.payload.Name).Return([]entity.Tiger{}, expectedErr)
		suite.mockTigerRepository.AssertNotCalled(t, "Insert", suite.ctx, suite.payloadInsert)

		result, err := suite.service.Create(suite.ctx, suite.payload)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.MessageResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on tiger already exist", func(t *testing.T) {
		suite := setupTest()
		respTiger := []entity.Tiger{
			{
				ID:          1,
				DateOfBirth: suite.payload.DateOfBirth.Format(time.DateOnly),
				LastLat:     40,
				LastLong:    40,
				LastSeen:    suite.service.fnTimeNow(),
				LastPhoto:   "public/no_image.svg",
				Name:        "tiger_new",
				CreatedAt:   suite.service.fnTimeNow(),
				UpdatedAt:   suite.service.fnTimeNow(),
			},
		}
		expectedErr := customerror.ErrorDuplicateTiger

		suite.mockTigerRepository.On("FindByName", suite.ctx, suite.payload.Name).Return(respTiger, nil)
		suite.mockTigerRepository.AssertNotCalled(t, "Insert", suite.ctx, suite.payloadInsert)

		result, err := suite.service.Create(suite.ctx, suite.payload)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.MessageResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on Insert failed", func(t *testing.T) {
		suite := setupTest()
		expectedErr := customerror.ErrorInternalServer

		suite.mockTigerRepository.On("FindByName", suite.ctx, suite.payload.Name).Return([]entity.Tiger{}, nil)
		suite.mockTigerRepository.On("Insert", suite.ctx, suite.payloadInsert).Return(expectedErr)

		result, err := suite.service.Create(suite.ctx, suite.payload)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.MessageResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})
}

func TestTigerService_ListSighting(t *testing.T) {
	type TestSuite struct {
		ctx                 context.Context
		mockTigerRepository *mockrepo.TigerRepositoryProvider
		service             *TigerService
	}

	setupTest := func() *TestSuite {
		timeNow, _ := time.Parse(time.RFC3339, "2024-01-01T12:33:12Z00:00")
		mockTigerRepository := &mockrepo.TigerRepositoryProvider{}
		service := NewTigerService(nil, mockTigerRepository)
		service.fnTimeNow = func() time.Time {
			return timeNow
		}

		return &TestSuite{
			ctx:                 context.Background(),
			mockTigerRepository: mockTigerRepository,
			service:             service,
		}
	}

	t.Run("success", func(t *testing.T) {
		suite := setupTest()
		payload := model.ListSightingRequest{
			PaginationRequest: model.PaginationRequest{
				Page:    1,
				PerPage: 3,
			},
			TigerID: 1,
		}
		respCount := uint64(10)
		respSighting := []entity.Sighting{
			{
				Username:  "user_1",
				TigerName: "tiger_1",
				Photo:     "public/no_image.svg",
				Lat:       50,
				Long:      50,
				CreatedAt: suite.service.fnTimeNow(),
			},
			{
				Username:  "user_1",
				TigerName: "tiger_1",
				Photo:     "public/no_image.svg",
				Lat:       40,
				Long:      40,
				CreatedAt: suite.service.fnTimeNow(),
			},
			{
				Username:  "user_1",
				TigerName: "tiger_1",
				Photo:     "public/no_image.svg",
				Lat:       45,
				Long:      45,
				CreatedAt: suite.service.fnTimeNow(),
			},
		}
		expectedResult := model.ListSightingResponse{
			Data: respSighting,
			Pagination: model.PaginationResponse{
				Page:      payload.Page,
				PerPage:   payload.PerPage,
				TotalPage: 4,
				TotalItem: respCount,
			},
		}

		suite.mockTigerRepository.On("CountSightingByTigerID", suite.ctx, payload.TigerID).Return(respCount, nil)
		suite.mockTigerRepository.On("FindSightingsByTigerIDWithPagination", suite.ctx, payload.TigerID, payload.PaginationRequest).Return(respSighting, nil)

		result, err := suite.service.ListSighting(suite.ctx, payload)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("success on page over total_page", func(t *testing.T) {
		suite := setupTest()
		payload := model.ListSightingRequest{
			PaginationRequest: model.PaginationRequest{
				Page:    5,
				PerPage: 3,
			},
			TigerID: 1,
		}
		respCount := uint64(10)
		respSighting := []entity.Sighting{}
		expectedResult := model.ListSightingResponse{
			Data: respSighting,
			Pagination: model.PaginationResponse{
				Page:      payload.Page,
				PerPage:   payload.PerPage,
				TotalPage: 4,
				TotalItem: respCount,
			},
		}

		suite.mockTigerRepository.On("CountSightingByTigerID", suite.ctx, payload.TigerID).Return(respCount, nil)
		suite.mockTigerRepository.On("FindSightingsByTigerIDWithPagination", suite.ctx, payload.TigerID, payload.PaginationRequest).Return(respSighting, nil)

		result, err := suite.service.ListSighting(suite.ctx, payload)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on CountSightingByTigerID failed", func(t *testing.T) {
		suite := setupTest()
		payload := model.ListSightingRequest{
			PaginationRequest: model.PaginationRequest{
				Page:    5,
				PerPage: 3,
			},
			TigerID: 1,
		}
		expectedErr := customerror.ErrorInternalServer

		suite.mockTigerRepository.On("CountSightingByTigerID", suite.ctx, payload.TigerID).Return(uint64(0), expectedErr)
		suite.mockTigerRepository.AssertNotCalled(t, "FindSightingsByTigerIDWithPagination", suite.ctx, payload.TigerID, payload.PaginationRequest)

		result, err := suite.service.ListSighting(suite.ctx, payload)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.ListSightingResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})

	t.Run("failed on FindSightingsByTigerIDWithPagination failed", func(t *testing.T) {
		suite := setupTest()
		payload := model.ListSightingRequest{
			PaginationRequest: model.PaginationRequest{
				Page:    1,
				PerPage: 3,
			},
			TigerID: 1,
		}
		expectedErr := customerror.ErrorInternalServer

		suite.mockTigerRepository.On("CountSightingByTigerID", suite.ctx, payload.TigerID).Return(uint64(10), nil)
		suite.mockTigerRepository.On("FindSightingsByTigerIDWithPagination", suite.ctx, payload.TigerID, payload.PaginationRequest).Return([]entity.Sighting{}, expectedErr)

		result, err := suite.service.ListSighting(suite.ctx, payload)

		assert.EqualError(t, expectedErr, err.Error())
		assert.Equal(t, model.ListSightingResponse{}, result)
		suite.mockTigerRepository.AssertExpectations(t)
	})
}
