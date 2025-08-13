package global

import (
	"go-simple-api/pkg/logger"
	"go-simple-api/settings"

	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
