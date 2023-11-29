package handler

import "go.uber.org/zap"

type Manager struct {
}

func NewManager(srv *service.Service, logger *zap.SugaredLogger) *Manager {
	return &Manager{}
}
