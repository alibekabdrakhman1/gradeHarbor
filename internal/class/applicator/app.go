package applicator

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/class/config"
	"go.uber.org/zap"
)

type App struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func New(logger *zap.SugaredLogger, cfg *config.Config) *App {
	return &App{
		config: cfg,
		logger: logger,
	}
}
