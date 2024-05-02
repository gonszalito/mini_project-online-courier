package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
)

func InitHTTPRoute(
	g *gin.Engine,
	mongodbClient mongodb.IMongoDB,
	ctx context.Context,
) {

	g.GET("/health-check", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{"status": "OK"})
	})

	InitOrderRoute("/order", ctx, g, mongodbClient)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Error: &model.ErrorLog{
				Message:       "Not Found",
				SystemMessage: "Not Found",
			},
		})
	})
}
