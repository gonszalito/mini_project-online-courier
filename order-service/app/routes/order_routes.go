package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/controllers"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/middlewares"
)

func InitOrderRoute(
	path string,
	ctx context.Context,
	g *gin.Engine,
	mongodb mongodb.IMongoDB,
) {
	ctrl := controllers.InitHTTPOrderController(mongodb, ctx)

	oauthMiddlewareGroup := g.Group("", middlewares.OauthMiddlware())

	oauthMiddlewareGroup.Use()

	userControllerGroup := oauthMiddlewareGroup.Group(path)
	userControllerGroup.Use()
	{
		userControllerGroup.POST("/create", ctrl.Create)
		userControllerGroup.GET("/user", ctrl.GetAllByUserID)
		userControllerGroup.GET("/pending", ctrl.GetAllPending)

		userControllerGroup.GET("/courier", ctrl.GetAllByCourierID)
		userControllerGroup.PATCH("/update", ctrl.Update)

	}
}
