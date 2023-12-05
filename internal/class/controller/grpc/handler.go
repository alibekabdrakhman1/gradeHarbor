package grpc

import (
	"context"
	"github.com/alibekabdrakhman1/gradeHarbor/pkg/enums"
	proto "github.com/alibekabdrakhman1/gradeHarbor/pkg/proto/class"
)

func (s *Server) GetMyUsers(ctx context.Context, in *proto.MyUsersRequest) (*proto.MyUsersResponse, error) {
	res, err := s.service.Class.GetMyStudents(ctx, uint(in.Id), in.Role)
	if err != nil {
		return nil, err
	}
	var students []uint32
	for _, v := range res {
		students = append(students, uint32(v))
	}
	var teachers []uint32
	if in.Role == enums.Student {
		res, err := s.service.Class.GetMyTeachers(ctx, uint(in.Id))
		if err != nil {
			return nil, err
		}
		for _, v := range res {
			teachers = append(teachers, uint32(v))
		}
	}

	return &proto.MyUsersResponse{
		Students: students,
		Teachers: teachers,
	}, nil
}

func (s *Server) GetClasses(ctx context.Context, in *proto.ClassRequest) (*proto.ClassResponse, error) {
	classes, err := s.service.Class.GetClassesByID(ctx, uint(in.Id), in.Role)
	if err != nil {
		return nil, err
	}

	var ans []*proto.Class
	for _, val := range classes {
		ans = append(ans, &proto.Class{
			Id:          uint32(val.ID),
			ClassCode:   val.ClassCode,
			ClassName:   val.ClassName,
			Description: val.Description,
			TeacherId:   uint32(val.TeacherID),
		})
	}

	return &proto.ClassResponse{Classes: ans}, nil
}

func (s *Server) GetGrades(ctx context.Context, in *proto.GradesRequest) (*proto.GradesResponse, error) {
	grades, err := s.service.Class.GetStudentGradesByID(ctx, uint(in.Id))
	if err != nil {
		return nil, err
	}

	var res proto.GradesResponse
	var gradesProto []*proto.Grades

	for _, internalGrade := range grades {
		gradeProto := &proto.Grades{
			ClassId:   uint32(internalGrade.ClassID),
			ClassCode: internalGrade.ClassCode,
			ClassName: internalGrade.ClassName,
			TeacherId: uint32(internalGrade.Teacher.ID),
			Students:  make([]*proto.GradeStudent, 0),
		}

		for _, internalStudent := range internalGrade.Students {
			gradeStudent := &proto.GradeStudent{
				Id:       uint32(internalStudent.ID),
				FullName: internalStudent.FullName,
				Grades:   make([]*proto.GradeResponse, 0),
			}

			for _, internalGradeResponse := range internalStudent.Grades {
				gradeResponse := &proto.GradeResponse{
					Grade: int32(internalGradeResponse.Grade),
					Week:  int32(internalGradeResponse.Week),
				}
				gradeStudent.Grades = append(gradeStudent.Grades, gradeResponse)
			}

			gradeProto.Students = append(gradeProto.Students, gradeStudent)
		}

		gradesProto = append(gradesProto, gradeProto)
	}

	res.Grades = gradesProto

	return &res, nil
}
