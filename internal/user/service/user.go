package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/model"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/storage"
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/transport"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enums"
	error2 "github.com/alibekabdrakhman1/gradeHarbor/pkg/error"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/utils"
	"go.uber.org/zap"
)

type UserService struct {
	repository         *storage.Repository
	logger             *zap.SugaredLogger
	classGrpcTransport *transport.ClassGrpcTransport
}

func NewUserService(r *storage.Repository, logger *zap.SugaredLogger, grpcTransport *transport.ClassGrpcTransport) *UserService {
	return &UserService{
		repository:         r,
		logger:             logger,
		classGrpcTransport: grpcTransport,
	}
}

func (s *UserService) Create(ctx context.Context, user model.User) (uint, error) {
	user.ParentID = 0
	user.IsConfirmed = false
	id, err := s.repository.User.Create(ctx, user)
	if err != nil {
		s.logger.Error(fmt.Sprintf("creating new user error: %v", err))
		return 0, fmt.Errorf("creating new user error: %v", err)
	}
	return id, nil
}

func (s *UserService) GetByID(ctx context.Context, userID uint) (*model.UserResponse, error) {
	err := s.checkPermission(ctx, userID, "")
	if err != nil {
		return nil, err
	}
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.repository.User.GetProfileById(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Errorf("GetProfileByID error: %v", err))
		return nil, fmt.Errorf("GetProfileByID error: %v", err)
	}
	return user, nil
}

func (s *UserService) GetByContext(ctx context.Context) (*model.UserResponse, error) {
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.User.GetById(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repository.User.GetByEmail(ctx, email)
}

func (s *UserService) Update(ctx context.Context, user model.User) (*model.User, error) {
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	oldUser, err := s.repository.User.GetById(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}

	if user.Email != oldUser.Email {
		return nil, errors.New("can not change email")
	}

	return s.repository.User.Update(ctx, user, id)
}

func (s *UserService) Delete(ctx context.Context) error {
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.repository.User.Delete(ctx, id)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *UserService) DeleteByID(ctx context.Context, userID uint) error {
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if role != enums.Admin {
		return errors.New("not permitted")
	}

	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return err
	}
	if user.Role == "admin" {
		s.logger.Error("cannot delete admin")
		return errors.New("cannot delete admin")
	}

	err = s.repository.User.Delete(ctx, userID)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *UserService) GetStudentTeachersByID(ctx context.Context, userID uint) ([]*model.UserResponse, error) {
	err := s.checkPermission(ctx, userID, enums.Student)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enums.Student {
		s.logger.Error(errors.New("user is not a student"))
		return nil, errors.New("user is not a student")
	}

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

func (s *UserService) GetStudentParentByID(ctx context.Context, userID uint) (*model.ParentResponse, error) {
	err := s.checkPermission(ctx, userID, enums.Student)
	if err != nil {
		return nil, err
	}
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enums.Student {
		s.logger.Error(errors.New("user is not a student"))
		return nil, errors.New("user is not a student")
	}

	parent, err := s.repository.User.GetStudentParentByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting student parent error: %v", err))
		return nil, fmt.Errorf("getting student parent error: %v", err)
	}

	return parent, nil
}

func (s *UserService) GetParentChildrenByID(ctx context.Context, userID uint) ([]*model.UserResponse, error) {
	err := s.checkPermission(ctx, userID, enums.Parent)
	if err != nil {
		return nil, err
	}
	id, err := utils.GetIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := s.repository.User.GetById(ctx, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("getting by id error: %v", err))
		return nil, err
	}
	if user.Role != enums.Parent {
		s.logger.Error(errors.New("user is not a parent"))
		return nil, errors.New("user is not a parent")
	}

	children, err := s.repository.User.GetParentChildrenByID(ctx, id)
	if err != nil {
		s.logger.Error(fmt.Errorf("getting parent children error: %v", err))
		return nil, fmt.Errorf("getting parent children error: %v", err)
	}

	return children, nil
}

func (s *UserService) checkPermission(ctx context.Context, id uint, role string) error {
	users, err := s.getMyUsersID(ctx)
	if err != nil {
		return err
	}
	if role == "" {
		for _, ids := range users {
			for _, val := range ids {
				if val == id {
					return nil
				}
			}
		}
		return error2.ErrNotPermitted
	}

	curr := users[role]
	for _, val := range curr {
		if val == id {
			return nil
		}
	}
	return error2.ErrNotPermitted
}

func (s *UserService) GetUserClassesByID(ctx context.Context, id uint) ([]*model.Class, error) {
	user, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.getClasses(ctx, id, user.Role)
}

func (s *UserService) GetStudentGradesByID(ctx context.Context, id uint) ([]*model.Grade, error) {
	user, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.Role == enums.Parent {
		return nil, errors.New("user role cannot be parent")
	}

	grades, err := s.getGrades(ctx, id)
	if err != nil {
		return nil, err
	}

	return grades, nil
}

func (s *UserService) getMyUsersID(ctx context.Context) (map[string][]uint, error) {
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
	var parents []uint
	for _, v := range users.Students {
		students = append(students, uint(v))
	}
	for _, v := range users.Teachers {
		teachers = append(teachers, uint(v))
	}
	for _, v := range users.Students {
		parent, err := s.GetStudentParentByID(ctx, uint(v))
		if err != nil {
			return nil, err
		}
		parents = append(parents, parent.ID)
	}
	m := make(map[string][]uint)
	m["student"] = students
	m["teacher"] = teachers
	m["parent"] = parents
	return m, nil
}

func (s *UserService) getGrades(ctx context.Context, id uint) ([]*model.Grade, error) {
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

func (s *UserService) getClasses(ctx context.Context, id uint, role string) ([]*model.Class, error) {
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
