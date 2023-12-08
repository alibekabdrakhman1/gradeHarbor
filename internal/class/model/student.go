package model

type ClassStudent struct {
	Id          string `json:"id"`
	ClassID     string `json:"class_id"`
	StudentID   uint   `db:"student_id" json:"student_id"`
	StudentName string `db:"student_name" json:"student_name"`
}
