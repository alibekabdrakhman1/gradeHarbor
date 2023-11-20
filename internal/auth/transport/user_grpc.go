package transport

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/auth/config"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserGrpcTransport struct {
	config config.UserGrpcTransport
	client proto.UserServiceClient
	logger *zap.SugaredLogger
}

func NewUserGrpcTransport(config config.UserGrpcTransport, logger *zap.SugaredLogger) *UserGrpcTransport {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.Dial(config.Port, opts...)
	if err != nil {
		logger.Errorf("grpc connect error: %v", err)
		log.Fatalf("error conn grpc: %s", err)
	}

	client := proto.NewUserServiceClient(conn)

	return &UserGrpcTransport{
		client: client,
		config: config,
		logger: logger,
	}
}

func (t *UserGrpcTransport) CreateUser(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.CreateUserResponse, error) {
	response, err := t.client.CreateUser(ctx, in)
	if err != nil {
		t.logger.Errorf("gprc CreateUser error: %v", err)
		return nil, fmt.Errorf("cannot CreateUser: %w", err)
	}
	if response == nil {
		t.logger.Errorf("grpc: not found")
		return nil, fmt.Errorf("not found")
	}
	return response, nil
}
