package model

type Class struct {
	ID          uint   `db:"id" json:"id"`
	ClassCode   string `db:"class_code" json:"class_code"`
	ClassName   string `db:"class_name" json:"class_name"`
	Description string `db:"description" json:"description"`
	TeacherID   uint   `db:"teacher_id" json:"teacher_id"`
	TeacherName string `db:"teacher_name" json:"teacher_name"`
}

type ClassResponse struct {
	ID          uint   `db:"id" json:"id"`
	ClassCode   string `db:"class_code" json:"class_code"`
	ClassName   string `db:"class_name" json:"class_name"`
	Description string `db:"description" json:"description"`
	Teacher     User   `json:"teacher"`
}

type ClassRequest struct {
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	Teacher     User   `json:"teacher"`
	Students    []User `json:"students"`
}

type User struct {
	ID       uint   `db:"student_id" json:"id"`
	FullName string `db:"student_name" json:"full_name"`
}

type ClassWithID struct {
	ID          uint   `db:"id" json:"id"`
	ClassCode   string `db:"class_code" json:"class_code"`
	ClassName   string `db:"class_name" json:"class_name"`
	Description string `db:"description" json:"description"`
	Teacher     User   `json:"teacher"`
	Students    []User `json:"students"`
}

type Relationships struct {
	ID        uint `db:"id" json:"id"`
	StudentID uint `db:"student_id" json:"student_id"`
	TeacherID uint `db:"teacher_id" json:"teacher_id"`
}
