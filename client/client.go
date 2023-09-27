package client

import (
	"context"
	"io"

	"github.com/ffauzann/authentication/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient interface {
	io.Closer
}

type Options struct {
	GrpcAddress  string
	Interceptors []grpc.UnaryClientInterceptor
}

func New(opts Options) (AuthClient, error) {
	conn, err := grpc.Dial(opts.GrpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(opts.Interceptors...),
	)
	if err != nil {
		return nil, err
	}
	return &authClient{
		conn:   conn,
		client: proto.NewAuthServiceClient(conn),
	}, nil
}

type authClient struct {
	conn   *grpc.ClientConn
	client proto.AuthServiceClient
}

func (c *authClient) Close() error {
	return c.conn.Close()
}

func (c *authClient) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	return c.client.Login(ctx, req)
}
