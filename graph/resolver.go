package graph

import (
	"github.com/go-playground/validator/v10"
	"github.com/vekaputra/tiger-kittens/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Validate     *validator.Validate
	TigerService service.TigerServiceProvider
	UserService  service.UserServiceProvider
}
