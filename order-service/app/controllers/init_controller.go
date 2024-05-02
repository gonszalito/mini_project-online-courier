package controllers

import (
	"context"

	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/repositories"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/usecases"
)

func InitHTTPOrderController(
	mongodb mongodb.IMongoDB,
	ctx context.Context,
) IOrderController {
	orderRepository := repositories.NewOrderRepository(mongodb)
	orderUseCase := usecases.NewOrderUseCase(orderRepository, mongodb, ctx)
	handler := NewOrderController(ctx, mongodb, orderUseCase, orderRepository)
	return handler
}
