package app

import (
	"github.com/vekaputra/tiger-kittens/internal/config"
	"github.com/vekaputra/tiger-kittens/internal/service"
)

type Service struct {
	UserService service.UserServiceProvider
}

func NewService(repo Repo, appConfig *config.Config) Service {
	userConfig := service.UserConfig{
		PrivateKey: appConfig.JWTConfig.PrivateKey,
	}

	return Service{
		UserService: service.NewUserService(userConfig, repo.UserRepo),
	}
}
