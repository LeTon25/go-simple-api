package routers_users

import (
	"go-simple-api/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()
	//public
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/register", userController.Register)
		userRouterPublic.GET("/detail/:id")
	}

	userRouterPrivate := Router.Group("/user/private")
	//userRouterPrivate.Use(Authen()) // Add middleware if needed
	//userRouterPrivate.Use(Permission()) // Placeholder for middleware, e.g., authentication
	{
		userRouterPrivate.GET("/info")
		userRouterPrivate.GET("/settings")
	}

}
