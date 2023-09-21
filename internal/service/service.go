package service

import (
	"context"

	"hangoutin/authentication/internal/model"
	"hangoutin/authentication/internal/repository"

	"go.uber.org/zap"
)

type Service interface {
	AuthService
}

type AuthService interface {
	IsUserExist(ctx context.Context, req *model.IsUserExistRequest) (res *model.IsUserExistResponse, err error)
	Register(ctx context.Context, req *model.RegisterRequest) (res *model.RegisterResponse, err error)
	Login(ctx context.Context, req *model.LoginRequest) (res *model.LoginResponse, err error)
}

type service struct {
	config     *model.AppConfig
	logger     *zap.Logger
	repository repositoryWrapper
}

type repositoryWrapper struct {
	db    repository.DBRepository
	redis repository.RedisRepository
}

func New(db repository.DBRepository, redis repository.RedisRepository, config *model.AppConfig, logger *zap.Logger) Service {
	return &service{
		config: config,
		logger: logger,
		repository: repositoryWrapper{
			db:    db,
			redis: redis,
		},
	}
}
