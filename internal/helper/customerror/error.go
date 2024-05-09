package customerror

import _const "github.com/vekaputra/tiger-kittens/internal/const"

var (
	ErrorDuplicateUser      = NewClientError(_const.ErrDuplicateUser, "email and username already exists")
	ErrorInvalidRequestBody = NewClientError(_const.ErrInvalidRequestBody, "invalid request, please check your request body")
	ErrorInvalidCredential  = NewClientError(_const.ErrInvalidCredential, "user not found")

	ErrorInternalServer = NewInternalServerError(_const.ErrInternalServer, "something went wrong")
)
