package model

import "time"

type ClassGrade struct {
	ID           uint      `json:"id"`
	ClassID      uint      `json:"class_id"`
	StudentID    uint      `json:"student_id"`
	Grade        int       `json:"grade"`
	Week         int       `json:"week"`
	LastModified time.Time `json:"last_modified"`
}

type GradesRequest struct {
	Grades []ClassGradeCreate `json:"grades"`
}

type ClassGradeCreate struct {
	StudentID uint `json:"student_id"`
	Grade     uint `json:"grade"`
	Week      int  `json:"week"`
}

type Grade struct {
	ClassCode string         `json:"class_code"`
	ClassName string         `json:"class_name"`
	Teacher   User           `json:"teacher_name"`
	Students  []GradeStudent `json:"students"`
}

type GradeStudent struct {
	ID       uint            `json:"id"`
	FullName string          `json:"full_name"`
	Grades   []GradeResponse `json:"grades"`
}

type GradeResponse struct {
	Grade        int       `json:"grade"`
	Week         int       `json:"week"`
	LastModified time.Time `json:"last_modified"`
}
