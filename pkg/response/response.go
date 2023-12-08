package response

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type IDResponse struct {
	ID uint `json:"id"`
}
