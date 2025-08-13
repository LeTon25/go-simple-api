package initalize

import (
	"go-simple-api/global"
	"go-simple-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Log)
}
