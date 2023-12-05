package service

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/transport"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"go.uber.org/zap"
)

func NewTeacherService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger, grpcTransport *transport.ClassGrpcTransport) *TeacherService {
	return &TeacherService{
		repository:         r,
		config:             cfg,
		logger:             logger,
		classGrpcTransport: grpcTransport,
	}
}

type TeacherService struct {
	repository         *storage.Repository
	config             *config.Config
	logger             *zap.SugaredLogger
	classGrpcTransport *transport.ClassGrpcTransport
}

func (s *TeacherService) GetStudents(ctx context.Context) ([]*model.UserResponse, error) {
	users, err := s.getMyUsersID(ctx)
	if err != nil {
		return nil, err
	}
	var students []*model.UserResponse

	for _, val := range users["student"] {
		student, err := s.repository.User.GetById(ctx, val)
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (s *TeacherService) getMyUsersID(ctx context.Context) (map[string][]uint, error) {
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	users, err := s.classGrpcTransport.GetMyUsers(ctx, &proto.MyUsersRequest{
		Id:   uint32(id),
		Role: role,
	})
	if err != nil {
		return nil, err
	}

	var students []uint
	var teachers []uint

	for _, v := range users.Students {
		students = append(students, uint(v))
	}
	for _, v := range users.Teachers {
		teachers = append(teachers, uint(v))
	}

	m := make(map[string][]uint)
	m["student"] = students
	m["teacher"] = teachers
	return m, nil
}
