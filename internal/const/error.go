package _const

type ErrorCode string

const (
	ErrDuplicateUser      ErrorCode = "duplicate_user"
	ErrInternalServer               = "internal_server_error"
	ErrInvalidRequestBody           = "invalid_request_body"
	ErrUnknownError                 = "unknown_error"
)
