package grpc

import (
	"context"
	"ewallet-ums/internal/domain/auth"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
)

type TokenValidationServer struct {
	pb.UnimplementedTokenValidationServiceServer
	Handler auth.ITokenValidationHandler
}

func (h *TokenValidationServer) ValidateToken(
	ctx context.Context,
	req *pb.ValidateTokenRequest,
) (*pb.ValidateTokenResponse, error) {
	return h.Handler.TokenValidation(ctx, req)
}
