package model

type ClassStudent struct {
	Id          string `json:"id"`
	ClassId     string `json:"class_id"`
	StudentId   uint   `json:"student_id"`
	StudentName string `json:"student_name"`
}
