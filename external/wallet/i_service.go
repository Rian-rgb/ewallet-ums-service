package wallet

import (
	"context"
)

type IService interface {
	CreateWallet(ctx context.Context, userID int) (*Wallet, error)
}
