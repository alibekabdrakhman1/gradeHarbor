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
func (s *Server) GetUserByEmail(ctx context.Context, in *proto.GetUserByEmailRequest) (*proto.GetUserByEmailResponse, error) {
	user, err := s.service.User.GetByEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, err
	}

	response := &proto.GetUserByEmailResponse{
		User: &proto.GetUser{
			Id:       uint32(user.ID),
			FullName: user.FullName,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role,
		},
	}
	return response, nil
}

func (s *Server) ConfirmUser(ctx context.Context, in *proto.ConfirmUserRequest) (*proto.ConfirmUserResponse, error) {
	err := s.service.User.ConfirmUser(ctx, in.GetEmail())
	if err != nil {
		return nil, err
	}

	response := &proto.ConfirmUserResponse{}
	return response, nil
}
