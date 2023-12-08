package transport

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClassGrpcTransport struct {
	config config.ClassGrpcTransport
	client proto.ClassServiceClient
	logger *zap.SugaredLogger
}

func NewClassGrpcTransport(config config.ClassGrpcTransport, logger *zap.SugaredLogger) *ClassGrpcTransport {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.Dial(config.Port, opts...)
	if err != nil {
		logger.Errorf("grpc connect error: %v", err)
		log.Fatalf("error conn grpc: %s", err)
	}
	fmt.Printf("connected grpc by port: %v", config.Port)

	client := proto.NewClassServiceClient(conn)

	return &ClassGrpcTransport{
		client: client,
		config: config,
		logger: logger,
	}
}

func (t *ClassGrpcTransport) GetMyUsers(ctx context.Context, in *proto.MyUsersRequest, opts ...grpc.CallOption) (*proto.MyUsersResponse, error) {
	users, err := t.client.GetMyUsers(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (t *ClassGrpcTransport) GetClasses(ctx context.Context, in *proto.ClassRequest, opts ...grpc.CallOption) (*proto.ClassResponse, error) {
	classes, err := t.client.GetClasses(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (t *ClassGrpcTransport) GetGrades(ctx context.Context, in *proto.GradesRequest, opts ...grpc.CallOption) (*proto.GradesResponse, error) {
	grades, err := t.client.GetGrades(ctx, in, opts...)
	if err != nil {
		fmt.Println("----")
		return nil, err
	}

	return grades, nil
}
