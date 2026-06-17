package grpc

import (
	"ewallet-ums/infra"
	"github.com/Rian-rgb/ewallet-common-lib/config"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	pb "github.com/Rian-rgb/ewallet-proto/gen/token_validation/v1"
	"google.golang.org/grpc"
	"net"
)

func NewGRPCServer(dependency *infra.Dependency) *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterTokenValidationServiceServer(
		grpcServer,
		&TokenValidationServer{
			Handler: dependency.TokenValidateAPI,
		},
	)

	return grpcServer
}

func ServeGRPC(dependency *infra.Dependency) {
	port := config.GetEnv("GRPC_PORT", "7000")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Error("failed to listen gRPC port: ", err)
	}

	grpcServer := NewGRPCServer(dependency)

	logger.Info("gRPC server listening on port %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("failed to serve gRPC server: ", err)
	}
}
