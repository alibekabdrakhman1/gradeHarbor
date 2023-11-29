package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enum"
	"go.uber.org/zap"
)

func NewClientService(r *storage.Repository, logger *zap.SugaredLogger) *AdminService {
	return &AdminService{
		repository: r,
		logger:     logger,
	}
}

type ClientService struct {
	repository *storage.Repository
	logger     *zap.SugaredLogger
}

func (s *ClientService) GetAllParents(ctx context.Context) ([]*model.ParentResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return nil, errors.New("not valid context userRole")
	}
	var parents []*model.ParentResponse
	var err error
	switch role.Role {
	case enum.Parent:
		return nil, errors.New("you are parent")
	case enum.Teacher:
		parents, err = s.repository.Teacher.GetAllParents(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("teacher's GetAllParents error: %v", err))
			return nil, err
		}
		return parents, nil
	case enum.Student:
		parents, err = s.repository.Student.GetAllParents(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("student's GetAllParents error: %v", err))
			return nil, err
		}
		return parents, nil
	}
	return parents, err
}

func (s *ClientService) GetParentByID(ctx context.Context, parentID uint) (*model.ParentResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return nil, errors.New("not valid context userRole")
	}
	var parent *model.ParentResponse
	var err error
	switch role.Role {
	case enum.Parent:
		return nil, errors.New("you are parent")
	case enum.Teacher:
		parent, err = s.repository.Teacher.GetParentByID(ctx, id.ID, parentID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("teacher's GetParentByID error: %v", err))
			return nil, err
		}
		return parent, nil
	case enum.Student:
		parent, err = s.repository.Student.GetParentByID(ctx, id.ID, parentID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("student's GetParentByID error: %v", err))
			return nil, err
		}
		return parent, nil
	}
	return parent, err
}

func (s *ClientService) GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return nil, errors.New("not valid context userRole")
	}

	var teachers []*model.UserResponse
	var err error

	switch role.Role {
	case enum.Parent:
		teachers, err = s.repository.Parent.GetAllTeachers(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("parent's GetAllTeachers error: %v", err))
			return nil, err
		}
		return teachers, nil
	case enum.Teacher:
		teachers, err = s.repository.Teacher.GetAllTeachers(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("teacher's GetAllTeachers error: %v", err))
			return nil, err
		}
		return teachers, nil
	case enum.Student:
		teachers, err = s.repository.Student.GetAllTeachers(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("student's GetAllTeachers error: %v", err))
			return nil, err
		}
		return teachers, nil
	}

	return teachers, err
}

func (s *ClientService) GetTeacherByID(ctx context.Context, teacherID uint) (*model.TeacherResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return nil, errors.New("not valid context userRole")
	}
	var teacher *model.TeacherResponse
	var err error
	switch role.Role {
	case enum.Parent:
		teacher, err = s.repository.Parent.GetTeacherByID(ctx, id.ID, teacherID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("parent's GetTeacherByID error: %v", err))
			return nil, err
		}
		return teacher, nil
	case enum.Teacher:
		teacher, err = s.repository.Teacher.GetTeacherByID(ctx, id.ID, teacherID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("teacher's GetTeacherByID error: %v", err))
			return nil, err
		}
		return teacher, nil
	case enum.Student:
		teacher, err = s.repository.Student.GetTeacherByID(ctx, id.ID, teacherID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("student's GetTeacherByID error: %v", err))
			return nil, err
		}
		return teacher, nil
	}
	return teacher, err
}

func (s *ClientService) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return nil, errors.New("not valid context userRole")
	}
	var students []*model.UserResponse
	var err error
	switch role.Role {
	case enum.Parent:
		students, err = s.repository.Parent.GetAllStudents(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("parent's GetAllStudents error: %v", err))
			return nil, err
		}
		return students, nil
	case enum.Teacher:
		students, err = s.repository.Teacher.GetAllStudents(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("teacher's GetAllStudents error: %v", err))
			return nil, err
		}
		return students, nil
	case enum.Student:
		students, err = s.repository.Student.GetAllStudents(ctx, id.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("student's GetAllStudents error: %v", err))
			return nil, err
		}
		return students, nil
	}
	return students, err
}

func (s *ClientService) GetStudentByID(ctx context.Context, studentID uint) (*model.StudentResponse, error) {
	id, ok := ctx.Value(model.ContextUserIDKey).(*model.ContextUserID)
	if !ok {
		s.logger.Error("not valid context userID")
		return nil, errors.New("not valid context userID")
	}
	role, ok := ctx.Value(model.ContextUserRoleKey).(*model.ContextUserRole)
	if !ok {
		s.logger.Error("not valid context userRole")
		return nil, errors.New("not valid context userRole")
	}
	var student *model.StudentResponse
	var err error
	switch role.Role {
	case enum.Parent:
		student, err = s.repository.Parent.GetStudentByID(ctx, id.ID, studentID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("parent's GetStudentByID error: %v", err))
			return nil, err
		}
		return student, nil
	case enum.Teacher:
		student, err = s.repository.Teacher.GetStudentByID(ctx, id.ID, studentID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("teacher's GetStudentByID error: %v", err))
			return nil, err
		}
		return student, nil
	case enum.Student:
		student, err = s.repository.Student.GetStudentByID(ctx, id.ID, studentID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("student's GetStudentByID error: %v", err))
			return nil, err
		}
		return student, nil
	}
	return student, err
}
