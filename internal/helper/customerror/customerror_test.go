package customerror

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomError(t *testing.T) {
	TestTooManyRequestsError := NewTooManyRequestError("too_many_requests", "too many requests")
	assert.Equal(t, CustomError{
		code:     "too_many_requests",
		message:  "too many requests",
		httpCode: http.StatusTooManyRequests,
	}, TestTooManyRequestsError)

	TesClientError := NewClientError("invalid_request_body", "Invalid Request, Please Check Your Request Body.")
	assert.Equal(t, CustomError{
		code:     "invalid_request_body",
		message:  "Invalid Request, Please Check Your Request Body.",
		httpCode: http.StatusBadRequest,
	}, TesClientError)

	TestUnauthorizedError := NewUnauthorizedError("invalid_client_credential", "Invalid Client Credentials")
	assert.Equal(t, CustomError{
		code:     "invalid_client_credential",
		message:  "Invalid Client Credentials",
		httpCode: http.StatusUnauthorized,
	}, TestUnauthorizedError)

	TestForbiddenError := NewForbiddenError("forbidden_error", "sorry, you don't have access to do this action")
	assert.Equal(t, CustomError{
		code:     "forbidden_error",
		message:  "sorry, you don't have access to do this action",
		httpCode: http.StatusForbidden,
	}, TestForbiddenError)

	TestNotFoundError := NewNotFoundError("not_found", "cannot find data")
	assert.Equal(t, CustomError{
		code:     "not_found",
		message:  "cannot find data",
		httpCode: http.StatusNotFound,
	}, TestNotFoundError)

	TestInternalServerError := NewInternalServerError("unexpected_error", "sorry, we have problem with our server. Please try again later")
	assert.Equal(t, CustomError{
		code:     "unexpected_error",
		message:  "sorry, we have problem with our server. Please try again later",
		httpCode: http.StatusInternalServerError,
	}, TestInternalServerError)
	assert.Equal(t, "unexpected_error", TestInternalServerError.Code())
	assert.Equal(t, "sorry, we have problem with our server. Please try again later", TestInternalServerError.Error())
	assert.Equal(t, http.StatusInternalServerError, TestInternalServerError.HTTPCode())
}
