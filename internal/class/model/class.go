package model

type Class struct {
	Id          uint   `json:"id"`
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	TeacherId   string `json:"teacher_id"`
	WeekNum     int    `json:"week_num"`
}

type ClassCreate struct {
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	TeacherID   string `json:"teacher_id"`
	WeekNum     int    `json:"week_num"`
	StudentsID  []uint `json:"students_id"`
}
