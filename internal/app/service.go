package app

import "github.com/vekaputra/tiger-kittens/internal/service"

type Service struct {
	UserService service.UserServiceProvider
}

func NewService(repo Repo) Service {
	return Service{
		UserService: service.NewUserService(repo.UserRepo),
	}
}
