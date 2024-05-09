package customerror

import (
	"errors"
	"net/http"

	_const "github.com/vekaputra/tiger-kittens/internal/const"
)

type CustomError struct {
	code     _const.ErrorCode
	message  string
	httpCode int
	metadata map[string]interface{}
}

func (ce CustomError) Error() string {
	return ce.message
}

func (ce CustomError) HTTPCode() int {
	return ce.httpCode
}

func (ce CustomError) Code() string {
	return string(ce.code)
}

func (ce CustomError) Equal(err error) bool {
	var e CustomError
	if errors.As(err, &e) {
		return ce.code == e.code && ce.httpCode == e.httpCode && ce.message == e.message
	}
	return false
}

func NewErrorWithCode(code _const.ErrorCode, message string, httpCode int) CustomError {
	return CustomError{code, message, httpCode, nil}
}

func NewClientError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusBadRequest)
}

func NewUnprocessableEntityError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusUnprocessableEntity)
}

func NewForbiddenError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusForbidden)
}

func NewUnauthorizedError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusUnauthorized)
}

func NewTooManyRequestError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusTooManyRequests)
}

func NewNotFoundError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusNotFound)
}

func NewInternalServerError(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusInternalServerError)
}

func NewUnsupportedMediaType(code _const.ErrorCode, message string) CustomError {
	return NewErrorWithCode(code, message, http.StatusUnsupportedMediaType)
}
