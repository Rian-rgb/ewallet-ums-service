package auth

import (
	"context"
	"ewallet-ums/internal/domain/user"
)

type ILoginService interface {
	Login(
		ctx context.Context,
		username string,
		password string,
	) (
		respUserEntity *user.Entity,
		respToken string,
		respRefreshToken string,
		err error,
	)
}
