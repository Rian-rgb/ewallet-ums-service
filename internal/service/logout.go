package service

import (
	"context"
	"ewallet-ums/internal/domain/user"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
)

type LogoutService struct {
	UserRepo  user.IRepository
	RedisRepo *redis.RedisRepository
}

func (svc *LogoutService) Logout(ctx context.Context, token string) error {

	userTokenKey := redis.UserTokenPrefix + token
	refreshTokenKey, err := svc.RedisRepo.Get(ctx, userTokenKey)
	if err != nil {

		logger.WithContext(ctx).Error("failed to get token in redis: ", err)
		return internalErrors.ErrInvalidToken
	}

	err = svc.RedisRepo.Delete(ctx, userTokenKey)
	if err != nil {

		logger.WithContext(ctx).Error("failed to delete token in redis: ", err)
		return internalErrors.ErrInvalidToken
	}

	err = svc.RedisRepo.Delete(ctx, refreshTokenKey)
	if err != nil {

		logger.WithContext(ctx).Error("failed to delete refresh token in redis: ", err)
		return internalErrors.ErrInvalidToken
	}

	return nil
}
