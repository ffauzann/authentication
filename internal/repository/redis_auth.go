package repository

import (
	"context"
	"fmt"
	"hangoutin/authentication/internal/constant"
	"hangoutin/authentication/internal/model"
)

func (r *redisRepository) RegisterUserDevice(ctx context.Context, deviceId string, token *model.Token) (err error) {
	key := fmt.Sprintf("%s:%s", constant.RedisPrefixGranted, deviceId)
	return r.redis.Set(ctx, key, token, 0).Err()
}
