package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/transport"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enums"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"go.uber.org/zap"
)

func NewAdminService(r *storage.Repository, cfg *config.Config, logger *zap.SugaredLogger, grpcTransport *transport.ClassGrpcTransport) *AdminService {
	return &AdminService{
		repository:         r,
		config:             cfg,
		logger:             logger,
		classGrpcTransport: grpcTransport,
	}
}

type AdminService struct {
	repository         *storage.Repository
	config             *config.Config
	logger             *zap.SugaredLogger
	classGrpcTransport *transport.ClassGrpcTransport
}

func (s *AdminService) GetAllTeachers(ctx context.Context) ([]*model.UserResponse, error) {
	teachers, err := s.repository.Admin.GetAllTeachers(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting all teachers error: %v", err)
	}
	return teachers, nil
}

func (s *AdminService) GetAllStudents(ctx context.Context) ([]*model.UserResponse, error) {
	students, err := s.repository.Admin.GetAllStudents(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting all students error: %v", err)
	}
	return students, nil
}

func (s *AdminService) GetAllParents(ctx context.Context) ([]*model.ParentResponse, error) {
	parents, err := s.repository.Admin.GetAllParents(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting all parents error: %v", err)
	}
	return parents, nil
}

func (s *AdminService) DeleteUserByID(ctx context.Context, id uint) error {
	err := s.repository.Admin.DeleteUserByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return fmt.Errorf("deleting user error: %v", err)
	}
	return nil
}

func (s *AdminService) PutParent(ctx context.Context, studentID uint, parentID uint) error {
	student, err := s.GetUserByID(ctx, studentID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting student by id error: %v", err))
		return err
	}
	if student.Role != enums.Student {
		return errors.New(fmt.Sprintf("user by %v is not student", studentID))
	}
	parent, err := s.GetUserByID(ctx, parentID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting parent by id error: %v", err))
		return err
	}
	if parent.Role != enums.Parent {
		return errors.New(fmt.Sprintf("user by %v is not parent", parentID))
	}
	err = s.repository.Admin.PutParent(ctx, studentID, parentID)
	if err != nil {
		s.logger.Error(err)
		return fmt.Errorf("putting parent for student error: %v", err)
	}
	return nil
}

func (s *AdminService) GetUserByID(ctx context.Context, id uint) (*model.UserResponse, error) {
	user, err := s.repository.Admin.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting user by id error: %v", err)
	}
	return user, nil
}

func (s *AdminService) GetStudentTeachersByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enums.Student {
		return nil, errors.New("this user is not a student")
	}

	users, err := s.getMyUsersID(ctx, id, user.Role)
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

func (s *AdminService) GetUserClassesByID(ctx context.Context, id uint) ([]*model.Class, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.getClasses(ctx, id, user.Role)
}

func (s *AdminService) GetStudentParentByID(ctx context.Context, id uint) (*model.ParentResponse, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enums.Student {
		return nil, errors.New("this user is not a student")
	}
	parent, err := s.repository.Admin.GetStudentParentByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting student parent by id error: %v", err)
	}
	return parent, nil
}

func (s *AdminService) GetParentChildrenByID(ctx context.Context, id uint) ([]*model.UserResponse, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enums.Parent {
		return nil, errors.New("this user is not a parent")
	}
	children, err := s.repository.Admin.GetParentChildrenByID(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("getting parent children by id error: %v", err)
	}
	return children, nil
}

func (s *AdminService) GetStudentGradesByID(ctx context.Context, id uint) ([]*model.Grade, error) {
	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.Role == enums.Parent {
		return nil, errors.New("user role cannot be parent")
	}

	return s.GetStudentGradesByID(ctx, id)
}

func (s *AdminService) CreateAdmin(ctx context.Context, user model.User) (uint, error) {
	user.IsConfirmed = true
	user.Role = "admin"
	user.Password, _ = utils.HashPassword(user.Password)
	return s.repository.User.Create(ctx, user)
}

func (s *AdminService) getMyUsersID(ctx context.Context, id uint, role string) (map[string][]uint, error) {
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

func (s *AdminService) getGrades(ctx context.Context, id uint) ([]*model.Grade, error) {
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

func (s *AdminService) getClasses(ctx context.Context, id uint, role string) ([]*model.Class, error) {
	classes, err := s.classGrpcTransport.GetClasses(ctx, &proto.ClassRequest{
		Id:   uint32(id),
		Role: role,
	})
	if err != nil {
		return nil, err
	}

	var res []*model.Class
	for _, val := range classes.Classes {
		res = append(res, &model.Class{
			Id:          uint(val.Id),
			ClassCode:   val.ClassCode,
			ClassName:   val.ClassName,
			Description: val.Description,
			TeacherId:   uint(val.TeacherId),
		})
	}

	return res, nil
}
