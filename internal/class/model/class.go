package model

type Class struct {
	ID          uint   `json:"id"`
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	TeacherID   User   `json:"teacher_id"`
}

type ClassRequest struct {
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	Teacher     User   `json:"teacher"`
	Students    []User `json:"students"`
}

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type ClassWithID struct {
	ID          uint   `json:"id"`
	ClassCode   string `json:"class_code"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
	Teacher     User   `json:"teacher"`
	Students    []User `json:"students"`
}
