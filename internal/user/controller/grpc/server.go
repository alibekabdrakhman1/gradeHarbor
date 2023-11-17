package grpc

import (
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	proto.UserServiceServer
	service    *service.Service
	config     *config.UserGrpcTransport
	grpcServer *grpc.Server
}

func NewServer(service *service.Service, config *config.UserGrpcTransport) *Server {
	s := grpc.NewServer()

	srv := &Server{service: service, config: config, grpcServer: s}

	proto.RegisterUserServiceServer(s, srv)
	reflection.Register(s)

	return srv
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %s", s.config.Port)
	}

	log.Printf("Starting grpc server: %s", s.config.Port)

	err = s.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
