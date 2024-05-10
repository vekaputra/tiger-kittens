package app

import (
	"github.com/vekaputra/tiger-kittens/internal/config"
	"github.com/vekaputra/tiger-kittens/internal/helper/mailqueue"
	"github.com/vekaputra/tiger-kittens/internal/service"
)

type Service struct {
	TigerService service.TigerServiceProvider
	UserService  service.UserServiceProvider
}

func NewService(repo Repo, appConfig *config.Config, mailQueue mailqueue.Provider) Service {
	userConfig := service.UserConfig{
		JWTConfig: appConfig.JWTConfig,
	}

	return Service{
		TigerService: service.NewTigerService(mailQueue, repo.TigerRepo),
		UserService:  service.NewUserService(userConfig, repo.UserRepo),
	}
}
