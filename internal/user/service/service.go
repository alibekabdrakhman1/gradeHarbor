package service

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
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

func NewManager(repo *storage.Repository, config *config.Config, logger *zap.SugaredLogger) *Service {
	userService := NewUserService(repo, logger)
	authService := NewAuthService(config.Auth)
	adminService := NewAdminService(repo, config, logger)
	teacherService := NewTeacherService(repo, config, logger)
	parentService := NewParentService(repo, config, logger)
	studentService := NewStudentService(repo, config, logger)
	return &Service{
		User:    userService,
		Auth:    authService,
		Admin:   adminService,
		Parent:  parentService,
		Student: studentService,
		Teacher: teacherService,
	}
}
