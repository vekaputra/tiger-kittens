package _const

type ErrorCode string

const (
	ErrDuplicateTiger     ErrorCode = "duplicate_tiger"
	ErrDuplicateUser                = "duplicate_user"
	ErrImageNotSupported            = "image_format_not_supported"
	ErrInternalServer               = "internal_server_error"
	ErrInvalidRequestBody           = "invalid_request_body"
	ErrInvalidCredential            = "invalid_credential"
	ErrInvalidAccessToken           = "invalid_access_token"
	ErrSightingWithin5KM            = "sighting_within_5_km"
	ErrTigerNotFound                = "tiger_not_found"
	ErrUnknownError                 = "unknown_error"
)
