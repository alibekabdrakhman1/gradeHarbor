package model

type Class struct {
	Id        string `json:"id"`
	ClassCode string `json:"class_code"`
	ClassName string `json:"class_name"`
	TeacherId string `json:"teacher_id"`
}
