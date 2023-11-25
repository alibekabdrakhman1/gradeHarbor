package handler

import "github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"

type AdminHandler struct {
	service service.Service
}

func NewAdminHandler(service service.Service) *AdminHandler {
	return &AdminHandler{service: service}
}
