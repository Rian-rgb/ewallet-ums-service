package service

import (
	"context"
	"ewallet-ums/internal/domain/user"
	internalErrors "ewallet-ums/internal/errors"
	"fmt"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/security"
)

type RefrshTokenService struct {
	UserRepo   user.IRepository
	JwtManager *security.JWTManager
	RedisRepo  *redis.RedisRepository
}

func (svc *RefrshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim security.Token) (token string, err error) {

	token, err = svc.JwtManager.GenerateToken(
		tokenClaim.UserID,
		tokenClaim.Username,
		tokenClaim.FullName,
		security.UserTokenDuration,
	)

	if err != nil {

		logger.WithContext(ctx).Error("failed to generate token: ", err)
		return "", internalErrors.ErrInternalServerError
	}

	refreshTokenKey := redis.RefreshTokenPrefix + refreshToken
	userTokenKey := redis.UserTokenPrefix + token
	err = svc.RedisRepo.Set(ctx, userTokenKey, refreshTokenKey, redis.UserTokenDuration)
	fmt.Println("test")
	if err != nil {

		logger.WithContext(ctx).Error("failed to add user token in redis: ", err)
		return "", internalErrors.ErrInternalServerError
	}

	return token, nil
}
