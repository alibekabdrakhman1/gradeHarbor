package model

type ParentResponse struct {
	ID          uint            `gorm:"primary_key;auto_increment" json:"id"`
	FullName    string          `gorm:"not null" json:"full_name"`
	Email       string          `gorm:"unique;not null" json:"email"`
	IsConfirmed bool            `json:"is_confirmed"`
	Children    []*UserResponse `json:"children"`
}

type UserResponse struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	FullName    string `gorm:"not null" json:"full_name"`
	Email       string `gorm:"unique;not null" json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
	ParentID    uint   `json:"parent_id"`
	Role        string `gorm:"not null" json:"role"`
}

type TeacherResponse struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	FullName    string `gorm:"not null" json:"full_name"`
	Email       string `gorm:"unique;not null" json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type StudentResponse struct {
	ID          uint           `gorm:"primary_key;auto_increment" json:"id"`
	FullName    string         `gorm:"not null" json:"full_name"`
	Email       string         `gorm:"unique;not null" json:"email"`
	IsConfirmed bool           `json:"is_confirmed"`
	ParentID    uint           `json:"parent_id"`
	Classes     []Class        `json:"classes"`
	Teachers    []UserResponse `json:"teachers"`
}
