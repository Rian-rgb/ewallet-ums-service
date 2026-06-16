package handler

import (
	"context"
	"ewallet-ums/internal/domain/auth"
	internalErrors "ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
)

type TokenValidationHandler struct {
	TokenValidationService auth.ITokenValidationService
}

func (s *TokenValidationHandler) TokenValidation(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	var (
		token = req.Token
	)

	if token == "" {
		logger.WithContext(ctx).Error("authorization header is empty")
		return &pb.ValidateTokenResponse{
			Message: internalErrors.ErrInvalidToken.Error(),
		}, nil
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &pb.ValidateTokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &pb.ValidateTokenResponse{
		Message: response.SuccessMessage,
		Data: &pb.User{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.FullName,
		},
	}, nil
}
