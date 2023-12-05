package model

var (
	ContextUserIDKey   = contextKey("id")
	ContextUserRoleKey = contextKey("role")
)

type User struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	FullName    string `gorm:"not null" json:"full_name"`
	Email       string `gorm:"unique;not null" json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
	ParentID    uint   `json:"parent_id"`
	Password    string `gorm:"not null" json:"password"`
	Role        string `gorm:"not null" json:"role"`
}

type UserRegister struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ContextUserID struct {
	ID uint `json:"user_id"`
}
type ContextUserRole struct {
	Role string `json:"role"`
}

type contextKey string

type ParentIDReq struct {
	ID uint `json:"parent_id"`
}
