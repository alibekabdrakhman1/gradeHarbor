package grpc

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/user"
)

func (s *Server) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user := model.User{
		FullName: in.GetUser().GetFullName(),
		Email:    in.GetUser().GetEmail(),
		Password: in.GetUser().GetPassword(),
		Role:     in.GetUser().GetRole(),
	}
	id, err := s.service.User.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{Id: uint32(id)}, nil
}
