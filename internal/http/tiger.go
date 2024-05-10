package http

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/file"
	"github.com/vekaputra/tiger-kittens/internal/helper/response"
	"github.com/vekaputra/tiger-kittens/internal/model"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

const (
	LastPhotoKey = "last_photo"
)

func (s *AppServer) CreateTiger(e echo.Context) error {
	ctx := e.Request().Context()

	var payload model.CreateTigerRequest
	if err := e.Bind(&payload); err != nil {
		log.Error().Err(err).Msg("failed bind payload")
		return response.SendResponseWithNativeError(e, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody))
	}

	dateOfBirth, err := time.Parse(time.DateOnly, e.FormValue("date_of_birth"))
	if err != nil {
		log.Error().Err(err).Msg("failed to parse date_of_birth")
		return response.SendResponseWithNativeError(e, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody))
	}
	payload.DateOfBirth = dateOfBirth

	if err = e.Validate(payload); err != nil {
		log.Error().Err(err).Msg("failed validate payload")
		return response.SendResponseWithNativeError(e, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody))
	}

	filepath, err := file.Save(
		e,
		LastPhotoKey,
		file.ResizeOption{
			Width:    _const.ResizeWidth,
			Height:   _const.ResizeHeight,
			IsResize: _const.IsResizeImage,
		},
	)
	if err != nil {
		return response.SendResponseWithNativeError(e, err)
	}
	payload.LastPhoto = filepath

	result, err := s.TigerService.Create(ctx, payload)
	if err != nil {
		return response.SendResponseWithNativeError(e, err)
	}

	return response.SendSuccessResponse(e, result)
}
