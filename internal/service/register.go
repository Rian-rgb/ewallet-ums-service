package service

import (
	"context"
	"ewallet-ums/external/wallet"
	"ewallet-ums/internal/domain/user"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo       user.IRepository
	ExternalWallet wallet.IService
}

func (svc *RegisterService) Register(ctx context.Context, user *user.Entity) (*user.Entity, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashPassword)

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
