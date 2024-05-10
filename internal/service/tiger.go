package service

import (
	"context"
	"time"

	"github.com/jftuga/geodist"
	"github.com/rs/zerolog/log"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
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
	ListSighting(ctx context.Context, page model.PaginationRequest) (model.ListSightingResponse, error)
}

type TigerService struct {
	TigerRepository pgsql.TigerRepositoryProvider
	UserRepository  pgsql.UserRepositoryProvider
	fnTimeNow       func() time.Time
}

func NewTigerService(tigerRepository pgsql.TigerRepositoryProvider, userRepository pgsql.UserRepositoryProvider) *TigerService {
	return &TigerService{
		TigerRepository: tigerRepository,
		UserRepository:  userRepository,
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

	return model.MessageResponse{
		Message:   _const.CreateSightingSuccessMessage,
		Timestamp: s.fnTimeNow().Format(time.RFC3339),
	}, nil
}

func (s *TigerService) ListSighting(ctx context.Context, page model.PaginationRequest) (model.ListSightingResponse, error) {
	count, err := s.TigerRepository.CountSighting(ctx)
	if err != nil {
		return model.ListSightingResponse{}, err
	}

	sightings, err := s.TigerRepository.FindSightingWithPagination(ctx, page, "created_at DESC")
	if err != nil {
		return model.ListSightingResponse{}, err
	}
	if len(sightings) == 0 {
		return model.ListSightingResponse{
			Data: []model.Sighting{},
			Pagination: model.PaginationResponse{
				Page:      page.Page,
				PerPage:   page.PerPage,
				TotalPage: pagination.TotalPage(count, page.PerPage),
				TotalItem: count,
			},
		}, nil
	}

	tigerMapByID, err := s.findTigersBySightingsAsMap(ctx, sightings)
	if err != nil {
		return model.ListSightingResponse{}, err
	}

	userMapByID, err := s.findUsersBySightingsAsMap(ctx, sightings)
	if err != nil {
		return model.ListSightingResponse{}, err
	}

	var result []model.Sighting
	for _, sighting := range sightings {
		result = append(result, model.Sighting{
			Username:  userMapByID[sighting.UserID].Username,
			TigerName: tigerMapByID[sighting.TigerID].Name,
			Photo:     sighting.Photo,
			Lat:       sighting.Lat,
			Long:      sighting.Long,
			CreatedAt: sighting.CreatedAt,
		})
	}

	return model.ListSightingResponse{
		Data: result,
		Pagination: model.PaginationResponse{
			Page:      page.Page,
			PerPage:   page.PerPage,
			TotalPage: pagination.TotalPage(count, page.PerPage),
			TotalItem: count,
		},
	}, nil
}

func (s *TigerService) findTigersBySightingsAsMap(ctx context.Context, sightings []entity.TigerSighting) (map[int]entity.Tiger, error) {
	var tigerIDs []int
	tigerMap := map[int]entity.Tiger{}
	for _, sighting := range sightings {
		if _, ok := tigerMap[sighting.TigerID]; !ok {
			tigerMap[sighting.TigerID] = entity.Tiger{}
			tigerIDs = append(tigerIDs, sighting.TigerID)
		}
	}

	tigers, err := s.TigerRepository.FindByIDs(ctx, tigerIDs)
	if err != nil {
		return nil, err
	}

	for _, tiger := range tigers {
		tigerMap[tiger.ID] = tiger
	}

	return tigerMap, nil
}

func (s *TigerService) findUsersBySightingsAsMap(ctx context.Context, sightings []entity.TigerSighting) (map[string]entity.User, error) {
	var userIDs []string
	userMap := map[string]entity.User{}
	for _, sighting := range sightings {
		if _, ok := userMap[sighting.UserID]; !ok {
			userMap[sighting.UserID] = entity.User{}
			userIDs = append(userIDs, sighting.UserID)
		}
	}

	users, err := s.UserRepository.FindByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userMap[user.ID] = user
	}

	return userMap, nil
}
