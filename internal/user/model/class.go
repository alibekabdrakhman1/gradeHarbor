package model

import "time"

type Class struct {
	Id          uint   `json:"id"`
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	TeacherId   uint   `json:"teacher_id"`
}

type Grade struct {
	ClassID   uint           `json:"class_id"`
	ClassCode string         `json:"class_code"`
	ClassName string         `json:"class_name"`
	TeacherID uint           `json:"teacher_id"`
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
