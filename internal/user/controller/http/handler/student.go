package handler

import "github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"

type StudentHandler struct {
	service service.Service
}

func NewStudentHandler(service service.Service) *StudentHandler {
	return &StudentHandler{service: service}
}
