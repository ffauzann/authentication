package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	// "hangoutin/common/interceptor"

	authInterceptor "hangoutin/common/interceptor/authentication"
	ctxTagsInterceptor "hangoutin/common/interceptor/ctxtags"
	logInterceptor "hangoutin/common/interceptor/logging"
	recoveryInterceptor "hangoutin/common/interceptor/recovery"

	// grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"hangoutin/authentication/internal/constant"
	delivery "hangoutin/authentication/internal/delivery/grpc"
	"hangoutin/authentication/internal/repository"
	"hangoutin/authentication/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	Address string
	Port    Port
	Logger  Logger
}

type Port struct {
	GRPC uint32
	HTTP uint32
}

func (c *Config) StartServer() {
	c.startGRPCServer()
}

func (c *Config) startGRPCServer() {
	addr := fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port.GRPC)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	interceptors := []grpc.UnaryServerInterceptor{
		authInterceptor.UnaryServerInterceptor(
			authInterceptor.WithExcludedMethods(c.App.Auth.ExcludedMethods...),
		),
		// interceptor.Authentication([]string{}),
		// interceptor.Logger(c.Server.Logger.Zap),
		// interceptor.ErrorHandler(constant.MapGRPCErrCodes),
		logInterceptor.UnaryServerInterceptor(
			c.Server.Logger.Zap,
			logInterceptor.WithErrorParser(constant.MapGRPCErrCodes),
		),
		ctxTagsInterceptor.UnaryServerInterceptor(),
		recoveryInterceptor.UnaryServerInterceptor(c.Server.Logger.Zap),
		// interceptor.Recovery(c.Server.Logger.Zap),
		grpc_ctxtags.UnaryServerInterceptor(),
		// grpc_recovery.UnaryServerInterceptor(),
	}
	opts := grpc.ChainUnaryInterceptor(interceptors...)
	grpcServer := grpc.NewServer(opts)

	// Init repo
	dbRepo := repository.NewDB(c.Database.SQL.DB, c.App, c.Server.Logger.Zap)
	redisRepo := repository.NewRedis(c.Cache.Redis.Client, c.App, c.Server.Logger.Zap)

	// Init service
	svc := service.New(dbRepo, redisRepo, c.App, c.Server.Logger.Zap)

	delivery.New(grpcServer, svc, c.Server.Logger.Zap)
	fmt.Printf("grpc started on %s\n", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
		return
	}

	go gracefullyStop(grpcServer)
}

func gracefullyStop(grpcServer *grpc.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("Gracefully shutdown...")
	grpcServer.GracefulStop()
}
