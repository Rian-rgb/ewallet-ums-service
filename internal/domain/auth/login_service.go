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
		userEntity *user.Entity,
		token string,
		refreshToken string,
		err error,
	)
}
