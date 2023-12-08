package model

import "time"

type ClassGrade struct {
	ID           uint      `db:"id"            json:"id"`
	ClassID      uint      `db:"class_id"      json:"class_id"`
	StudentID    uint      `db:"student_id"    json:"student_id"`
	Grade        int       `db:"grade"         json:"grade"`
	Week         int       `db:"week"          json:"week"`
	LastModified time.Time `db:"last_modified" json:"last_modified"`
}

type GradesRequest struct {
	ID     uint               `db:"id"     json:"class_id"`
	Grades []ClassGradeCreate `db:"grades" json:"grades"`
}

type ClassGradeCreate struct {
	StudentID uint `db:"student_id" json:"student_id"`
	Grade     uint `db:"grade"      json:"grade"`
	Week      int  `db:"week"       json:"week"`
}

type Grade struct {
	ClassID   uint           `db:"class_id"     json:"class_id"`
	ClassCode string         `db:"class_code"   json:"class_code"`
	ClassName string         `db:"class_name"   json:"class_name"`
	Teacher   User           `db:"teacher_name" json:"teacher_name"`
	Students  []GradeStudent `json:"students"`
}

type GradeStudent struct {
	ID       uint            `db:"id"        json:"id"`
	FullName string          `db:"full_name" json:"full_name"`
	Grades   []GradeResponse `json:"grades"`
}

type GradeResponse struct {
	Grade        int       `db:"grade"         json:"grade"`
	Week         int       `db:"week"          json:"week"`
	LastModified time.Time `db:"last_modified" json:"last_modified"`
}
