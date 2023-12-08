package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type JwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Message struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

type MessageRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
