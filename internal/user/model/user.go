package model

type User struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	FullName    string `gorm:"not null" json:"full_name"`
	Email       string `gorm:"unique;not null" json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
	ParentID    string `json:"parent_id"`
	Password    string `gorm:"not null" json:"password"`
	Role        string `gorm:"not null" json:"role"`
}

type UserRegister struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
