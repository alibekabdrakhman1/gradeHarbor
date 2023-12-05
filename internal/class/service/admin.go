package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage"
	"go.uber.org/zap"
	"sort"
)

func NewAdminService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *AdminService {
	return &AdminService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type AdminService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *AdminService) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	return s.repository.Admin.GetClassGradesByID(ctx, id)
}

func (s *AdminService) CreateClass(ctx context.Context, class model.ClassRequest) (uint, error) {
	seen := make(map[uint]bool)
	for _, item := range class.Students {
		if seen[item.ID] {
			return 0, errors.New(fmt.Sprintf("student has a duplicate: %v", item.ID))
		}
		seen[item.ID] = true
	}
	sort.Slice(class.Students, func(i, j int) bool {
		return class.Students[i].FullName < class.Students[j].FullName
	})
	fmt.Println(class)
	return s.repository.Admin.CreateClass(ctx, class)
}

func (s *AdminService) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	return s.repository.Admin.GetAllClasses(ctx)
}

func (s *AdminService) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	return s.repository.Admin.GetClassByID(ctx, id)
}

func (s *AdminService) UpdateClassByID(ctx context.Context, id uint, class model.ClassRequest) (*model.ClassWithID, error) {
	return s.repository.Admin.UpdateClassByID(ctx, id, class)
}

func (s *AdminService) DeleteClassByID(ctx context.Context, id uint) error {
	return s.repository.Admin.DeleteClassByID(ctx, id)
}

func (s *AdminService) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	return s.repository.Admin.GetClassStudentsByID(ctx, id)
}

func (s *AdminService) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	return s.repository.Admin.GetClassTeacherByID(ctx, id)
}

func (s *AdminService) GetStudentGradesByID(ctx context.Context, studentID uint) ([]*model.Grade, error) {
	return s.repository.Class.GetStudentGradesByID(ctx, studentID)
}
