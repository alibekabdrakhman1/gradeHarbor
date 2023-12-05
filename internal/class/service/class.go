package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enums"
	error2 "github.com/alibekabdrakhman1/gradeHarbor/pkg/error"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"go.uber.org/zap"
)

func NewClassService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger) *ClassService {
	return &ClassService{
		repository: r,
		config:     cfg,
		logger:     logger,
	}
}

type ClassService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *zap.SugaredLogger
}

func (s *ClassService) GetClassesByID(ctx context.Context, userID uint, role string) ([]*model.Class, error) {
	switch role {
	case enums.Parent:
		return nil, errors.New("parent does not have any classes")
	case enums.Teacher:
		return s.repository.Class.GetClassesForTeacher(ctx, userID)
	case enums.Student:
		return s.repository.Class.GetClassesForStudent(ctx, userID)
	}
	return nil, errors.New("role is not correct")
}

func (s *ClassService) GetAllClasses(ctx context.Context) ([]*model.Class, error) {
	userID, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	switch role {
	case enums.Parent:
		return nil, errors.New("parent does not have any classes")
	case enums.Teacher:
		return s.repository.Class.GetClassesForTeacher(ctx, userID)
	case enums.Student:
		return s.repository.Class.GetClassesForStudent(ctx, userID)
	}
	return nil, errors.New("role is not correct")
}

func (s *ClassService) GetClassByID(ctx context.Context, id uint) (*model.ClassWithID, error) {
	userID, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	err = s.checkPermission(ctx, role, userID, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return s.repository.Class.GetClassByID(ctx, id)
}

func (s *ClassService) GetClassStudentsByID(ctx context.Context, id uint) ([]*model.User, error) {
	userID, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	err = s.checkPermission(ctx, role, userID, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return s.repository.Class.GetClassStudentsByID(ctx, id)
}

func (s *ClassService) GetClassGradesByID(ctx context.Context, id uint) (*model.Grade, error) {
	userID, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	err = s.checkPermission(ctx, role, userID, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	switch role {
	case enums.Teacher:
		return s.repository.Class.GetClassGradesByIDForTeacher(ctx, id, userID)
	case enums.Student:
		return s.repository.Class.GetClassGradesByIDForStudent(ctx, id, userID)
	}
	return nil, error2.ErrNotPermitted
}

func (s *ClassService) PutClassGradesByID(ctx context.Context, id uint, grades model.GradesRequest) error {
	userID, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	err = s.checkPermission(ctx, role, userID, id)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return s.repository.Class.PutClassGradesByID(ctx, id, grades)
}

func (s *ClassService) GetClassTeacherByID(ctx context.Context, id uint) (*model.User, error) {
	userID, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	err = s.checkPermission(ctx, role, userID, id)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return s.repository.Class.GetClassTeacherByID(ctx, id)
}

func (s *ClassService) GetStudentGradesByID(ctx context.Context, studentID uint) ([]*model.Grade, error) {
	return s.repository.Class.GetStudentGradesByID(ctx, studentID)
}

func (s *ClassService) GetMyStudents(ctx context.Context, userID uint, role string) ([]uint, error) {
	if role == enums.Teacher {
		return s.repository.Class.GetMyStudentsForTeacher(ctx, userID)
	}
	return s.repository.Class.GetMyStudentsForStudent(ctx, userID)
}

func (s *ClassService) GetMyTeachers(ctx context.Context, userID uint) ([]uint, error) {
	return s.repository.Class.GetMyTeachers(ctx, userID)
}

func (s *ClassService) checkPermission(ctx context.Context, role string, userID uint, classID uint) error {
	if role == enums.Teacher {
		teacher, err := s.GetClassTeacherByID(ctx, classID)
		if err != nil {
			return fmt.Errorf("getting class teacher error: %v", err)
		}

		if teacher.ID != userID {
			return error2.ErrNotPermitted
		}
		return nil
	} else {
		students, err := s.GetClassStudentsByID(ctx, classID)
		if err != nil {
			return fmt.Errorf("getting class students error: %v", err)
		}

		for _, student := range students {
			if student.ID == userID {
				return nil
			}
		}
		return error2.ErrNotPermitted
	}
}
