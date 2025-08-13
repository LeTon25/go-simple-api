package routers_admin

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *AdminRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public
	userRouterPublic := Router.Group("/admin")
	{
		userRouterPublic.POST("/login")
	}

	userRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(Authen()) // Add middleware if needed
	//userRouterPrivate.Use(Permission()) // Placeholder for middleware, e.g., authentication
	{
		userRouterPrivate.POST("/active_user")
	}

}
