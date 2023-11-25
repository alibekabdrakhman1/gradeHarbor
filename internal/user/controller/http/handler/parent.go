package handler

import "github.com/alibekabdrakhman1/gradeHarbor/internal/user/service"

type ParentHandler struct {
	service service.Service
}

func NewParentHandler(service service.Service) *ParentHandler {
	return &ParentHandler{service: service}
}
