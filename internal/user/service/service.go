package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/transport"
	"go.uber.org/zap"
)

type Service struct {
	User    IUserService
	Auth    IAuthService
	Admin   IAdminService
	Parent  IParentService
	Student IStudentService
	Teacher ITeacherService
}

func NewManager(repo *storage.Repository, config *config.Config, logger *zap.SugaredLogger, grpcTransport *transport.ClassGrpcTransport) *Service {
	userService := NewUserService(repo, logger, grpcTransport)
	authService := NewAuthService(config.Auth)
	adminService := NewAdminService(repo, config, logger, grpcTransport)
	teacherService := NewTeacherService(repo, config, logger, grpcTransport)
	parentService := NewParentService(repo, config, logger)
	studentService := NewStudentService(repo, config, logger, grpcTransport)
	return &Service{
		User:    userService,
		Auth:    authService,
		Admin:   adminService,
		Parent:  parentService,
		Student: studentService,
		Teacher: teacherService,
	}
}
