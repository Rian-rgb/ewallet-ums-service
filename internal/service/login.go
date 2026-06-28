package service

import (
	"context"
	"ewallet-ums/internal/domain/user"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LoginService struct {
	UserRepo   user.IRepository
	JwtManager *security.JWTManager
	RedisRepo  *redis.RedisRepository
}

func (svc *LoginService) Login(
	ctx context.Context,
	username string,
	password string,
) (
	userEntity *user.Entity,
	token string,
	refreshToken string,
	err error,
) {

	userEntity, err = svc.UserRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", "", internalErrors.ErrUserNotFound
		}

		logger.WithContext(ctx).Error("failed to find by username: ", err)
		return nil, "", "", internalErrors.ErrInternalServerError
	}

	err = security.VerifyPassword(userEntity.Password, password)
	if err != nil {
		return nil, "", "", internalErrors.ErrInvalidPassword
	}

	token, err = svc.JwtManager.GenerateToken(
		userEntity.ID,
		userEntity.Username,
		userEntity.FullName,
		security.UserTokenDuration,
	)

	if err != nil {

		logger.WithContext(ctx).Error("failed to generate token: ", err)
		return nil, "", "", internalErrors.ErrInternalServerError
	}

	refreshToken, err = svc.JwtManager.GenerateToken(
		userEntity.ID,
		userEntity.Username,
		userEntity.FullName,
		security.RefreshTokenDuration,
	)

	if err != nil {

		logger.WithContext(ctx).Error("failed to generate refresh token: ", err)
		return nil, "", "", internalErrors.ErrInternalServerError
	}

	refreshTokenKey := redis.RefreshTokenPrefix + refreshToken
	err = svc.RedisRepo.Set(ctx, refreshTokenKey, userEntity.ID, redis.RefreshTokenDuration)
	if err != nil {

		logger.WithContext(ctx).Error("failed to add refresh token in redis: ", err)
		return nil, "", "", internalErrors.ErrInternalServerError
	}

	userTokenKey := redis.UserTokenPrefix + token
	err = svc.RedisRepo.Set(ctx, userTokenKey, refreshTokenKey, redis.UserTokenDuration)
	if err != nil {

		logger.WithContext(ctx).Error("failed to add user token in redis: ", err)
		return nil, "", "", internalErrors.ErrInternalServerError
	}

	return userEntity, token, refreshToken, nil
}
