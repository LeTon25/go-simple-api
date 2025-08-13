package initalize

import (
	"go-simple-api/global"
	"go-simple-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware
	r.Use() // logging
	r.Use() // cross
	r.Use() // limit

	manageRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("api/v1")
	{
		MainGroup.GET("/check_status")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.AdminRouter.InitUserRouter(MainGroup)
		manageRouter.UserRouter.InitUserRouter(MainGroup)
	}

	return r
}
