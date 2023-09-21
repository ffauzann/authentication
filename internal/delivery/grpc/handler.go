package grpc

import (
	"hangoutin/authentication/internal/service"
	"hangoutin/authentication/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type srv struct {
	proto.UnimplementedAuthServiceServer
	service service.AuthService
	logger  *zap.Logger
}

func New(server *grpc.Server, authSrv service.AuthService, logger *zap.Logger) {
	srv := srv{
		service: authSrv,
		logger:  logger,
	}
	proto.RegisterAuthServiceServer(server, &srv)
	reflection.Register(server)
}
