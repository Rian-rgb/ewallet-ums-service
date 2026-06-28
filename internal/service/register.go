package service

import (
	"context"
	"ewallet-ums/external/wallet"
	"ewallet-ums/internal/domain/user"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/security"
)

type RegisterService struct {
	UserRepo       user.IRepository
	ExternalWallet wallet.IService
}

func (svc *RegisterService) Register(ctx context.Context, user *user.Entity) (*user.Entity, error) {
	hashPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPassword

	err = svc.UserRepo.Save(user)
	if err != nil {
		logger.WithContext(ctx).Error("failed to save user: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	logger.WithContext(ctx).Error("calling ewallet-wallet")
	_, err = svc.ExternalWallet.CreateWallet(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	resp := user
	resp.Password = ""
	return resp, nil
}
