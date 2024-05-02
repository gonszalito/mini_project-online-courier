package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/middlewares"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/routes"
)

func MainHTTPHandler(
	mongodbClient mongodb.IMongoDB,
	ctx context.Context,
) {

	g := gin.Default()
	g.Use(middlewares.CORSMiddleware(), middlewares.JSONMiddleware(), middlewares.RequestIdMiddleware())

	routes.InitHTTPRoute(g, mongodbClient, ctx)

	addr := fmt.Sprintf(":%s", os.Getenv("MAIN_PORT"))

	g.Run(addr)
}
