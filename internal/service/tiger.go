package service

import (
	"context"
	"time"

	"github.com/vekaputra/tiger-kittens/internal/helper/pagination"

	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	"github.com/vekaputra/tiger-kittens/internal/repository/pgsql"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

//go:generate mockery --name=TigerServiceProvider --outpkg=mock --output=./mock
type TigerServiceProvider interface {
	Create(ctx context.Context, payload model.CreateTigerRequest) (model.MessageResponse, error)
	List(ctx context.Context, page model.PaginationRequest) (model.ListTigerResponse, error)
}

type TigerService struct {
	TigerRepository pgsql.TigerRepositoryProvider
	fnTimeNow       func() time.Time
}

func NewTigerService(tigerRepository pgsql.TigerRepositoryProvider) *TigerService {
	return &TigerService{
		TigerRepository: tigerRepository,
		fnTimeNow:       time.Now,
	}
}

func (s *TigerService) Create(ctx context.Context, payload model.CreateTigerRequest) (model.MessageResponse, error) {
	tigers, err := s.TigerRepository.FindByName(ctx, payload.Name)
	if err != nil {
		return model.MessageResponse{}, err
	}
	if len(tigers) != 0 {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(customerror.ErrorDuplicateTiger)
	}

	err = s.TigerRepository.Insert(ctx, entity.Tiger{
		DateOfBirth: payload.DateOfBirth,
		LastLat:     payload.LastLat,
		LastLong:    payload.LastLong,
		LastPhoto:   payload.LastPhoto,
		LastSeen:    payload.LastSeen,
		Name:        payload.Name,
	})
	if err != nil {
		return model.MessageResponse{}, err
	}

	return model.MessageResponse{
		Message:   _const.CreateTigerSuccessMessage,
		Timestamp: s.fnTimeNow().Format(time.RFC3339),
	}, nil
}

func (s *TigerService) List(ctx context.Context, page model.PaginationRequest) (model.ListTigerResponse, error) {
	count, err := s.TigerRepository.Count(ctx)
	if err != nil {
		return model.ListTigerResponse{}, err
	}

	tigers, err := s.TigerRepository.FindWithPagination(ctx, page)
	if err != nil {
		return model.ListTigerResponse{}, err
	}

	return model.ListTigerResponse{
		Data: tigers,
		Pagination: model.PaginationResponse{
			Page:      page.Page,
			PerPage:   page.PerPage,
			TotalPage: pagination.TotalPage(count, page.PerPage),
			TotalItem: count,
		},
	}, nil
}
