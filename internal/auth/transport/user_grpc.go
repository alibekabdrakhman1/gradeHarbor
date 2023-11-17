package transport

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserGrpcTransport struct {
	config config.UserGrpcTransport
	client proto.UserServiceClient
}

func NewUserGrpcTransport(config config.UserGrpcTransport) *UserGrpcTransport {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.Dial(config.Port, opts...)
	if err != nil {
		log.Fatalf("error conn grpc: %s", err)
	}

	client := proto.NewUserServiceClient(conn)

	return &UserGrpcTransport{
		client: client,
		config: config,
	}
}

func (t *UserGrpcTransport) CreateUser(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.CreateUserResponse, error) {
	response, err := t.client.CreateUser(ctx, in)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("cannot CreateUser: %w", err)
	}
	if response == nil {
		return nil, fmt.Errorf("not found")
	}
	return response, nil
}
