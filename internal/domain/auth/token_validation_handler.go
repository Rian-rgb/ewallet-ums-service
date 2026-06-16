package auth

import (
	"context"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
)

type ITokenValidationHandler interface {
	TokenValidation(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error)
}
