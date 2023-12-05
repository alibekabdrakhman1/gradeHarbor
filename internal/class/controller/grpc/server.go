package grpc

import (
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/service"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	proto.ClassServiceServer
	service    *service.Service
	config     *config.ClassGrpcTransport
	grpcServer *grpc.Server
}

func NewServer(service *service.Service, config *config.ClassGrpcTransport) *Server {
	s := grpc.NewServer()

	srv := &Server{service: service, config: config, grpcServer: s}

	proto.RegisterClassServiceServer(s, srv)
	reflection.Register(s)

	return srv
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %s", s.config.Port)
	}

	log.Printf("Starting grpc server: %s", s.config.Port)

	go func() {
		err := s.grpcServer.Serve(listener)
		if err != nil {

		}
	}()

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
