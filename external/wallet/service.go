package wallet

import (
	"bytes"
	"context"
	"encoding/json"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"io"
	"net/http"
)

type ExtWallet struct{}

func (*ExtWallet) CreateWallet(ctx context.Context, userID int) (*Wallet, error) {
	var (
		walletHost = config.GetEnv("WALLET_HOST", "")
		urlAPI     = config.GetEnv("WALLET_ENDPOINT_CREATE", "")
	)

	req := Wallet{UserID: userID}
	payload, err := json.Marshal(req)
	if err != nil {
		logger.WithContext(ctx).Error("failed to prepare create wallet request: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	url := walletHost + urlAPI
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		logger.WithContext(ctx).Error("failed to create http request: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		logger.WithContext(ctx).Error("failed to send http request: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.WithContext(ctx).Error("failed to disconnect http: ", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		logger.WithContext(ctx).Error("failed to get response create wallet service: ", resp.StatusCode)
		return nil, internalErrors.ErrInternalServerError
	}

	result := Wallet{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logger.WithContext(ctx).Error("failed to read respose body: ", err)
		return nil, internalErrors.ErrInternalServerError
	}

	return &result, nil
}
