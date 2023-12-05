package service

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/transport"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"go.uber.org/zap"
)

func NewStudentService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger, grpcTransport *transport.ClassGrpcTransport) *StudentService {
	return &StudentService{
		repository:         r,
		config:             cfg,
		logger:             logger,
		classGrpcTransport: grpcTransport,
	}
}

type StudentService struct {
	repository         *storage.Repository
	config             *config.Config
	logger             *zap.SugaredLogger
	classGrpcTransport *transport.ClassGrpcTransport
}

func (s *StudentService) GetGroupmates(ctx context.Context) ([]*model.UserResponse, error) {
	users, err := s.getMyUsersID(ctx)
	if err != nil {
		return nil, err
	}
	var groupmates []*model.UserResponse

	for _, val := range users["student"] {
		groupmate, err := s.repository.User.GetById(ctx, val)
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
		groupmates = append(groupmates, groupmate)
	}

	return groupmates, nil
}

func (s *StudentService) GetParent(ctx context.Context) (*model.ParentResponse, error) {
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	parent, err := s.repository.Student.GetParent(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student parent error: %v", err))
		return nil, fmt.Errorf("getting student parent error: %v", err)
	}

	return parent, nil
}

func (s *StudentService) GetTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	users, err := s.getMyUsersID(ctx)
	if err != nil {
		return nil, err
	}
	var teachers []*model.UserResponse

	for _, val := range users["teacher"] {
		teacher, err := s.repository.User.GetById(ctx, val)
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (s *StudentService) GetGrades(ctx context.Context) ([]*model.Grade, error) {
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	grades, err := s.classGrpcTransport.GetGrades(ctx, &proto.GradesRequest{Id: uint32(id)})
	if err != nil {
		return nil, err
	}

	var res []*model.Grade

	for _, internalGrade := range grades.Grades {
		gradeProto := &model.Grade{
			ClassID:   uint(internalGrade.ClassId),
			ClassCode: internalGrade.ClassCode,
			ClassName: internalGrade.ClassName,
			TeacherID: uint(internalGrade.TeacherId),
			Students:  make([]model.GradeStudent, 0),
		}

		for _, internalStudent := range internalGrade.Students {
			gradeStudent := model.GradeStudent{
				ID:       uint(internalStudent.Id),
				FullName: internalStudent.FullName,
				Grades:   make([]model.GradeResponse, 0),
			}

			for _, internalGradeResponse := range internalStudent.Grades {
				gradeResponse := model.GradeResponse{
					Grade: int(internalGradeResponse.Grade),
					Week:  int(internalGradeResponse.Week),
				}
				gradeStudent.Grades = append(gradeStudent.Grades, gradeResponse)
			}

			gradeProto.Students = append(gradeProto.Students, gradeStudent)
		}

		res = append(res, gradeProto)
	}

	return res, nil
}

func (s *StudentService) getMyUsersID(ctx context.Context) (map[string][]uint, error) {
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
