package customerror

import _const "github.com/vekaputra/tiger-kittens/internal/const"

var (
	ErrorDuplicateTiger     = NewClientError(_const.ErrDuplicateTiger, "tiger already exists")
	ErrorDuplicateUser      = NewClientError(_const.ErrDuplicateUser, "email and username already exists")
	ErrorImageNotSupported  = NewClientError(_const.ErrImageNotSupported, "image uploaded not supported, please only upload 'jpeg' or 'png' image")
	ErrorInvalidRequestBody = NewClientError(_const.ErrInvalidRequestBody, "invalid request, please check your request body")
	ErrorInvalidCredential  = NewClientError(_const.ErrInvalidCredential, "user not found")

	ErrorInternalServer = NewInternalServerError(_const.ErrInternalServer, "something went wrong")
)
