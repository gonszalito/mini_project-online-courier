package grpcclient

import (
	"context"

	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/repositories"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/usecases"
)

func InitGrpcOauthClient(mongodb mongodb.IMongoDB, ctx context.Context) *OAuthGrpcClient {
	userRepository := repositories.NewUserRepository(mongodb)
	// userUseCase := usecases.NewUserUseCase(userRepository, mongodb, ctx)
	oauthUseCase := usecases.InitOAuthUseCase(userRepository)
	handler := InitOAuthGrpcClient(oauthUseCase)

	return handler
}
