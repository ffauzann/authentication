package repository

import (
	"context"

	"github.com/ffauzann/authentication/internal/constant"
	"github.com/ffauzann/authentication/internal/model"

	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func NewDB(db *sqlx.DB, config *model.AppConfig, logger *zap.Logger) DBRepository {
	return &dbRepository{
		db: db,
		common: common{
			config: config,
			logger: logger,
		},
	}
}

func NewRedis(client *redis.Client, config *model.AppConfig, logger *zap.Logger) RedisRepository {
	return &redisRepository{
		redis: client,
		common: common{
			config: config,
			logger: logger,
		},
	}
}

type DBRepository interface {
	// User-related
	CreateUser(ctx context.Context, user *model.User) error
	IsUserExist(ctx context.Context, userIdType constant.UserIdType, userIdVal string) (isExist bool, err error)
	GetUserByOneOfIdentifier(ctx context.Context, val string) (user *model.User, err error)

	// Role-related
	BatchAssignRoles(ctx context.Context, userId uint64, roleIds []uint8) error

	// Device-related
	RegisterUserDevice(ctx context.Context, userDevice *model.UserDevice) error
}

type RedisRepository interface {
	RegisterUserDevice(ctx context.Context, deviceId string, token *model.Token) error
}

type common struct {
	config *model.AppConfig
	logger *zap.Logger
}

type dbRepository struct {
	db *sqlx.DB
	common
}

type redisRepository struct {
	redis *redis.Client
	common
}
