package service

import (
	"context"
	"crypto/rsa"
	"time"

	"github.com/vekaputra/tiger-kittens/internal/helper/jwt"

	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/hash"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	"github.com/vekaputra/tiger-kittens/internal/repository/pgsql"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

type UserConfig struct {
	PrivateKey *rsa.PrivateKey
}

//go:generate mockery --name=UserServiceProvider --outpkg=mock --output=./mock
type UserServiceProvider interface {
	Register(ctx context.Context, payload model.RegisterUserRequest) (model.MessageResponse, error)
	Login(ctx context.Context, payload model.LoginUserRequest) (model.LoginUserResponse, error)
}

type UserService struct {
	Config         UserConfig
	UserRepository pgsql.UserRepositoryProvider
	fnTimeNow      func() time.Time
}

func NewUserService(config UserConfig, userRepository pgsql.UserRepositoryProvider) *UserService {
	return &UserService{
		Config:         config,
		UserRepository: userRepository,
		fnTimeNow:      time.Now,
	}
}

func (s *UserService) Register(ctx context.Context, payload model.RegisterUserRequest) (model.MessageResponse, error) {
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

func (s *UserService) Login(ctx context.Context, payload model.LoginUserRequest) (model.LoginUserResponse, error) {
	users, err := s.UserRepository.FindByEmailOrUsername(ctx, payload.Username, payload.Username)
	if err != nil {
		return model.LoginUserResponse{}, err
	}
	if len(users) == 0 {
		return model.LoginUserResponse{}, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidCredential)
	}

	user := users[0]
	if !hash.CheckBCrypt(payload.Password, user.Password) {
		return model.LoginUserResponse{}, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidCredential)
	}

	accessToken, err := jwt.GenerateAccessToken(s.Config.PrivateKey, user)
	if err != nil {
		return model.LoginUserResponse{}, err
	}

	return model.LoginUserResponse{
		AccessToken: accessToken,
		Timestamp:   s.fnTimeNow().Format(time.RFC3339),
	}, nil
}
