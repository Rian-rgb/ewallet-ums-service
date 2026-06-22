package auth

import (
	"context"
	"github.com/Rian-rgb/ewallet-common-lib/security"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim security.Token) (token string, err error)
}
