package auth

import (
	"context"
	"github.com/Rian-rgb/ewallet-common-lib/security"
)

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*security.ClaimToken, error)
}
