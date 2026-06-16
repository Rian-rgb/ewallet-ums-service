package auth

import (
	"context"
	"ewallet-ums/internal/domain/user"
)

type IRegisterService interface {
	Register(ctx context.Context, user *user.Entity) (*user.Entity, error)
}
