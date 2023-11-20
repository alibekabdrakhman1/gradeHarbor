package model

type Class struct {
	Id          uint   `json:"id"`
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	TeacherId   string `json:"teacher_id"`
}
