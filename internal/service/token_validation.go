package service

import (
	"context"
	"ewallet-ums/internal/domain/user"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"time"
)

type TokenValidationService struct {
	UserRepo   user.IRepository
	JwtManager *security.JWTManager
	RedisRepo  *redis.RedisRepository
}

func (svc *TokenValidationService) TokenValidation(ctx context.Context, token string) (*security.ClaimToken, error) {
	var (
		claimToken *security.ClaimToken
		err        error
	)

	claimToken, err = svc.JwtManager.ValidateToken(token)
	if err != nil {

		logger.WithContext(ctx).Error("failed to validate token: ", err)
		return nil, internalErrors.ErrInvalidToken
	}

	expTime, err := claimToken.GetExpirationTime()
	if err != nil || time.Now().After(expTime.Time) {
		logger.WithContext(ctx).Error("token has expired, expired at: ", claimToken.ExpiresAt)
		return nil, internalErrors.ErrTokenExpired
	}

	userTokenKey := redis.UserTokenPrefix + claimToken.ID
	exists, err := svc.RedisRepo.Exists(ctx, userTokenKey)
	if err != nil {

		logger.WithContext(ctx).Error("failed to get token from redis: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	if !exists {
		logger.WithContext(ctx).Error("token no exists in redis")
		return nil, internalErrors.ErrInvalidToken
	}

	return claimToken, nil
}
