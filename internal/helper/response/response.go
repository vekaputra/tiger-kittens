package response

import (
	"errors"
	"net/http"
	"time"

	_const "github.com/vekaputra/tiger-kittens/internal/const"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Timestamp        string `json:"timestamp"`
}

func SendResponseWithNativeError(e echo.Context, val error) error {
	log.Error().Err(val).Msg("an error occurred")
	now := time.Now().Format(time.RFC3339)

	var customError customerror.CustomError
	if !errors.As(pkgerr.ErrCause(val), &customError) {
		return e.JSON(http.StatusInternalServerError, Error{
			Error:            _const.ErrUnknownError,
			ErrorDescription: val.Error(),
			Timestamp:        now,
		})
	}
	return e.JSON(customError.HTTPCode(), Error{
		Error:            customError.Code(),
		ErrorDescription: customError.Error(),
		Timestamp:        now,
	})
}

func SendSuccessResponse(e echo.Context, val interface{}) error {
	return e.JSON(http.StatusOK, val)
}
