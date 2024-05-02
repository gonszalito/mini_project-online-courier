package controllers

import (
	"context"

	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/repositories"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/usecases"
)

func InitHTTPUserController(
	mongodb mongodb.IMongoDB,
	ctx context.Context,
) IUserController {
	userRepository := repositories.NewUserRepository(mongodb)
	userUseCase := usecases.NewUserUseCase(userRepository, mongodb, ctx)
	handler := NewUserController(ctx, mongodb, userUseCase, userRepository)

	return handler
}
