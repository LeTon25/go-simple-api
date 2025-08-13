package global

import (
	"go-simple-api/pkg/logger"
	"go-simple-api/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)
