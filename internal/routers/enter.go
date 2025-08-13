package routers

import (
	routers_admin "go-simple-api/internal/routers/admin"
	routers_users "go-simple-api/internal/routers/users"
)

type RouterGroup struct {
	User  routers_users.UserRouterGroup
	Admin routers_admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
