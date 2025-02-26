package middleware

import (
	"go-dvm/config"
	"go-dvm/pkg/logger"

	"github.com/labstack/echo/v4"
)

type MiddlewareManager struct {
	cfg    *config.Config
	logger *logger.Logger
	echo   *echo.Echo
}

func NewMiddlewareManager(cfg *config.Config, logger *logger.Logger, echo *echo.Echo) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:    cfg,
		logger: logger,
		echo:   echo,
	}
}
