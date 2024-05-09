package http

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/response"
	"github.com/vekaputra/tiger-kittens/internal/model"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

func (s *AppServer) RegisterUser(e echo.Context) error {
	ctx := e.Request().Context()

	var payload model.RegisterUserRequest
	if err := e.Bind(&payload); err != nil {
		log.Error().Err(err).Msg("failed bind payload")
		return response.SendResponseWithNativeError(e, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody))
	}
	if err := e.Validate(payload); err != nil {
		log.Error().Err(err).Msg("failed validate payload")
		return response.SendResponseWithNativeError(e, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidRequestBody))
	}

	result, err := s.UserService.Register(ctx, payload)
	if err != nil {
		return response.SendResponseWithNativeError(e, err)
	}

	return response.SendSuccessResponse(e, result)
}
