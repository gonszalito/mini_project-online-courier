package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/controllers"
)

func InitUserRoute(
	path string,
	ctx context.Context,
	g *gin.Engine,
	mongodb mongodb.IMongoDB,
) {
	ctrl := controllers.InitHTTPUserController(mongodb, ctx)

	// oauthMiddlewareGroup := g.Group("", authMiddleware)
	// oauthMiddlewareGroup.Use()

	userControllerGroup := g.Group(path)
	{
		userControllerGroup.POST("/signup", ctrl.SignUp)
		userControllerGroup.POST("/login", ctrl.Login)
		userControllerGroup.GET("/:username", ctrl.GetProfile)

	}
}
