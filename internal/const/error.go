package _const

type ErrorCode string

const (
	ErrDuplicateTiger     ErrorCode = "duplicate_tiger"
	ErrDuplicateUser                = "duplicate_user"
	ErrImageNotSupported            = "image_format_not_supported"
	ErrInternalServer               = "internal_server_error"
	ErrInvalidRequestBody           = "invalid_request_body"
	ErrInvalidCredential            = "invalid_credential"
	ErrUnknownError                 = "unknown_error"
)
