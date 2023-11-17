package http

import (
	"github.com/alibekabdrakhman1/gradeHarbor/internal/user/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config *config.Config
	App    *echo.Echo
}
