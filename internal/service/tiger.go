package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jftuga/geodist"
	"github.com/rs/zerolog/log"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/mailqueue"
	"github.com/vekaputra/tiger-kittens/internal/helper/pagination"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	"github.com/vekaputra/tiger-kittens/internal/repository/pgsql"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

//go:generate mockery --name=TigerServiceProvider --outpkg=mock --output=./mock
type TigerServiceProvider interface {
	Create(ctx context.Context, payload model.CreateTigerRequest) (model.MessageResponse, error)
	List(ctx context.Context, page model.PaginationRequest) (model.ListTigerResponse, error)
	CreateSighting(ctx context.Context, payload model.CreateSightingRequest) (model.MessageResponse, error)
	ListSighting(ctx context.Context, page model.ListSightingRequest) (model.ListSightingResponse, error)
}

type TigerService struct {
	MailQueue       mailqueue.Provider
	TigerRepository pgsql.TigerRepositoryProvider
	fnTimeNow       func() time.Time
}

func NewTigerService(mailQueue mailqueue.Provider, tigerRepository pgsql.TigerRepositoryProvider) *TigerService {
	return &TigerService{
		MailQueue:       mailQueue,
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
		DateOfBirth: payload.DateOfBirth.Format(time.DateOnly),
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

func (s *TigerService) CreateSighting(ctx context.Context, payload model.CreateSightingRequest) (model.MessageResponse, error) {
	var err, errTx error
	tx, err := s.TigerRepository.BeginTx(ctx)
	if err != nil {
		return model.MessageResponse{}, err
	}
	defer func() {
		errCloseTx := s.TigerRepository.CloseTx(tx, errTx)
		if errCloseTx != nil {
			log.Error().Err(errCloseTx).Msg("failed to close tx")
		}
	}()

	tiger, errTx := s.TigerRepository.TxFindByID(ctx, tx, payload.TigerID)
	if errTx != nil {
		return model.MessageResponse{}, errTx
	}
	if tiger == nil {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(customerror.ErrorTigerNotFound)
	}

	tigerCoord := geodist.Coord{
		Lat: tiger.LastLat,
		Lon: tiger.LastLong,
	}
	sightingCoord := geodist.Coord{
		Lat: payload.Lat,
		Lon: payload.Long,
	}
	_, kmDist := geodist.HaversineDistance(tigerCoord, sightingCoord)
	if kmDist < 5 {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(customerror.ErrorSightingWithin5KM)
	}

	lastSeen := s.fnTimeNow()
	errTx = s.TigerRepository.TxInsertSighting(ctx, tx, entity.TigerSighting{
		UserID:    payload.UserID,
		TigerID:   payload.TigerID,
		Photo:     payload.Photo,
		Lat:       payload.Lat,
		Long:      payload.Long,
		CreatedAt: lastSeen,
	})
	if errTx != nil {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(errTx)
	}

	tiger.LastLat = payload.Lat
	tiger.LastLong = payload.Long
	tiger.LastPhoto = payload.Photo
	tiger.LastSeen = lastSeen

	errTx = s.TigerRepository.TxUpdate(ctx, tx, *tiger)
	if errTx != nil {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(errTx)
	}

	emails, errTx := s.TigerRepository.TxFindSightingUploaderEmailsByTigerID(ctx, tx, tiger.ID)
	if errTx != nil {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(errTx)
	}

	title := fmt.Sprintf("New sightings of tiger (%s)", tiger.Name)
	body := fmt.Sprintf(
		"Hi, we just found a new sightings of tiger (%s) at lat: %.7f, long: %.7f. It's last seen around %s",
		tiger.Name,
		tiger.LastLat,
		tiger.LastLong,
		tiger.LastSeen.Format(time.RFC3339),
	)
	for _, email := range emails {
		go s.MailQueue.Add(mailqueue.SendEmailJob{
			DestinationEmail: email,
			Title:            title,
			Body:             body,
		})
	}

	return model.MessageResponse{
		Message:   _const.CreateSightingSuccessMessage,
		Timestamp: s.fnTimeNow().Format(time.RFC3339),
	}, nil
}

func (s *TigerService) ListSighting(ctx context.Context, payload model.ListSightingRequest) (model.ListSightingResponse, error) {
	count, err := s.TigerRepository.CountSighting(ctx)
	if err != nil {
		return model.ListSightingResponse{}, err
	}

	sightings, err := s.TigerRepository.FindSightingsByTigerIDWithPagination(ctx, payload.TigerID, payload.PaginationRequest)
	if err != nil {
		return model.ListSightingResponse{}, err
	}

	return model.ListSightingResponse{
		Data: sightings,
		Pagination: model.PaginationResponse{
			Page:      payload.PaginationRequest.Page,
			PerPage:   payload.PaginationRequest.PerPage,
			TotalPage: pagination.TotalPage(count, payload.PaginationRequest.PerPage),
			TotalItem: count,
		},
	}, nil
}
