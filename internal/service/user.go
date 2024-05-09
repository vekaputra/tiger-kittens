package service

import (
	"context"
	"time"

	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"

	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/hash"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	"github.com/vekaputra/tiger-kittens/internal/repository/pgsql"
)

//go:generate mockery --name=UserServiceProvider --outpkg=mock --output=./mock
type UserServiceProvider interface {
	RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (model.MessageResponse, error)
}

type UserService struct {
	UserRepository pgsql.UserRepositoryProvider
	fnTimeNow      func() time.Time
}

func NewUserService(userRepository pgsql.UserRepositoryProvider) *UserService {
	return &UserService{
		UserRepository: userRepository,
		fnTimeNow:      time.Now,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, payload model.RegisterUserRequest) (model.MessageResponse, error) {
	users, err := s.UserRepository.FindByEmailOrUsername(ctx, payload.Email, payload.Username)
	if err != nil {
		return model.MessageResponse{}, err
	}
	if len(users) != 0 {
		return model.MessageResponse{}, pkgerr.ErrWithStackTrace(customerror.ErrorDuplicateUser)
	}

	passwordHash, err := hash.BCrypt(payload.Password)
	if err != nil {
		return model.MessageResponse{}, err
	}

	err = s.UserRepository.Insert(ctx, entity.User{
		Email:    payload.Email,
		Password: passwordHash,
		Username: payload.Username,
	})
	if err != nil {
		return model.MessageResponse{}, err
	}

	return model.MessageResponse{
		Message:   _const.RegisterSuccessMessage,
		Timestamp: s.fnTimeNow().Format(time.RFC3339),
	}, nil
}
