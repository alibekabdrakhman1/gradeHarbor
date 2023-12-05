package model

var (
	ContextUserIDKey   = contextKey("id")
	ContextUserRoleKey = contextKey("role")
)

type ContextUserID struct {
	ID uint `json:"user_id"`
}
type ContextUserRole struct {
	Role string `json:"role"`
}

type contextKey string
