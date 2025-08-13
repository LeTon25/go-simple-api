//go:build wireinject

package wire

import (
	"go-simple-api/internal/controller"
	"go-simple-api/internal/repo"
	"go-simple-api/internal/services"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		services.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
