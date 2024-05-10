package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"fmt"
	intctx "github.com/vekaputra/tiger-kittens/internal/helper/context"
	"strconv"
	"time"

	_ "github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/graph/model"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/file"
	"github.com/vekaputra/tiger-kittens/internal/helper/pagination"
	intmodel "github.com/vekaputra/tiger-kittens/internal/model"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

// CreateTiger is the resolver for the createTiger field.
func (r *mutationResolver) CreateTiger(ctx context.Context, input model.CreateTiger) (*model.Message, error) {
	var payload intmodel.CreateTigerRequest

	payload.LastLat = input.LastLat
	payload.LastLong = input.LastLong
	payload.LastSeen = input.LastSeen
	payload.Name = input.Name

	dateOfBirth, err := time.Parse(time.DateOnly, input.DateOfBirth)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse date_of_birth")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody)
	}
	payload.DateOfBirth = dateOfBirth

	if err = r.Validate.Struct(payload); err != nil {
		log.Error().Err(err).Msg("failed validate payload")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody)
	}

	filepath, err := file.SaveGQL(
		input.LastPhoto.File,
		file.ResizeOption{
			Width:    _const.ResizeWidth,
			Height:   _const.ResizeHeight,
			IsResize: _const.IsResizeImage,
		},
	)
	if err != nil {
		return nil, err
	}
	payload.LastPhoto = filepath

	result, err := r.TigerService.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &model.Message{
		Message:   result.Message,
		Timestamp: result.Timestamp,
	}, nil
}

// CreateSighting is the resolver for the createSighting field.
func (r *mutationResolver) CreateSighting(ctx context.Context, input model.CreateSighting) (*model.Message, error) {
	var payload intmodel.CreateSightingRequest

	tigerID, err := strconv.Atoi(input.TigerID)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse tigerID")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody)
	}
	payload.TigerID = tigerID
	payload.UserID = intctx.GetUser(ctx)
	payload.Lat = input.Lat
	payload.Long = input.Long

	if err = r.Validate.Struct(payload); err != nil {
		log.Error().Err(err).Msg("failed validate payload")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody)
	}

	filepath, err := file.SaveGQL(
		input.Photo.File,
		file.ResizeOption{
			Width:    _const.ResizeWidth,
			Height:   _const.ResizeHeight,
			IsResize: _const.IsResizeImage,
		},
	)
	if err != nil {
		return nil, err
	}
	payload.Photo = filepath

	result, err := r.TigerService.CreateSighting(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &model.Message{
		Message:   result.Message,
		Timestamp: result.Timestamp,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.Login, error) {
	var payload intmodel.LoginUserRequest

	payload.Password = input.Password
	payload.Username = input.Username

	if err := r.Validate.Struct(payload); err != nil {
		log.Error().Err(err).Msg("failed validate payload")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody)
	}

	result, err := r.UserService.Login(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &model.Login{
		AccessToken: result.AccessToken,
		Timestamp:   result.Timestamp,
	}, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterUser) (*model.Message, error) {
	var payload intmodel.RegisterUserRequest

	payload.Email = input.Email
	payload.Password = input.Password
	payload.Username = input.Username

	if err := r.Validate.Struct(payload); err != nil {
		log.Error().Err(err).Msg("failed validate payload")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody)
	}

	result, err := r.UserService.Register(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &model.Message{
		Message:   result.Message,
		Timestamp: result.Timestamp,
	}, nil
}

// Tigers is the resolver for the tigers field.
func (r *queryResolver) Tigers(ctx context.Context, input *model.PaginationInput) (*model.ListTiger, error) {
	var page intmodel.PaginationRequest

	if input != nil {
		page.Page = uint64(input.Page)
		page.PerPage = uint64(input.PerPage)
	}

	result, err := r.TigerService.List(ctx, pagination.DefaultPagination(page))
	if err != nil {
		return nil, err
	}

	var data []*model.Tiger
	for _, rs := range result.Data {
		data = append(data, &model.Tiger{
			ID:          fmt.Sprint(rs.ID),
			DateOfBirth: rs.DateOfBirth,
			LastLat:     rs.LastLat,
			LastLong:    rs.LastLong,
			LastSeen:    rs.LastSeen,
			LastPhoto:   rs.LastPhoto,
			Name:        &rs.Name,
			CreatedAt:   rs.CreatedAt,
			UpdatedAt:   rs.UpdatedAt,
		})
	}
	return &model.ListTiger{
		Data: data,
		Pagination: &model.Pagination{
			Page:      int(result.Pagination.Page),
			PerPage:   int(result.Pagination.PerPage),
			TotalPage: int(result.Pagination.TotalPage),
			TotalItem: int(result.Pagination.TotalItem),
		},
	}, nil
}

// TigerSightings is the resolver for the tigerSightings field.
func (r *queryResolver) TigerSightings(ctx context.Context, input *model.PaginationInput) (*model.ListSighting, error) {
	var page intmodel.PaginationRequest

	if input != nil {
		page.Page = uint64(input.Page)
		page.PerPage = uint64(input.PerPage)
	}

	result, err := r.TigerService.ListSighting(ctx, pagination.DefaultPagination(page))
	if err != nil {
		return nil, err
	}

	var data []*model.Sighting
	for _, rs := range result.Data {
		data = append(data, &model.Sighting{
			UploadedBy: rs.Username,
			TigerName:  rs.TigerName,
			Photo:      rs.Photo,
			Lat:        rs.Lat,
			Long:       rs.Long,
			CreatedAt:  rs.CreatedAt,
		})
	}
	return &model.ListSighting{
		Data: data,
		Pagination: &model.Pagination{
			Page:      int(result.Pagination.Page),
			PerPage:   int(result.Pagination.PerPage),
			TotalPage: int(result.Pagination.TotalPage),
			TotalItem: int(result.Pagination.TotalItem),
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
