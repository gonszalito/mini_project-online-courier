package usecases

import (
	"context"

	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/models"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/repositories"
)

type IOrderUseCase interface {
	CreateOrder(request *models.Order) (*models.Order, *model.ErrorLog)
	UpdateOrder(request *models.Order) (*models.Order, *model.ErrorLog)
	GetAllByUserID(userID string) ([]*models.Order, *model.ErrorLog)
	GetAllByCourierID(courierID string) ([]*models.Order, *model.ErrorLog)
	GetByID(id string) (*models.Order, *model.ErrorLog)
	GetAllByUsername(username string) ([]*models.Order, *model.ErrorLog)
	GetAllByCourierUsername(username string) ([]*models.Order, *model.ErrorLog)
	GetAllPending() ([]*models.Order, *model.ErrorLog)
}

type OrderUseCase struct {
	orderRepository repositories.IOrderRepository
	mongod          mongodb.IMongoDB
	ctx             context.Context
}

func NewOrderUseCase(
	orderRepository repositories.IOrderRepository,
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IOrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
		mongod:          mongod,
		ctx:             ctx,
	}
}

func (uc *OrderUseCase) GetAllPending() ([]*models.Order, *model.ErrorLog) {
	result := make(chan models.OrdersChan)
	go uc.orderRepository.GetAll(uc.ctx, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	// Filter orders where status is "pending"
	var pendingOrders []*models.Order
	for _, order := range response.Orders {
		if order.Status == "pending" {
			pendingOrders = append(pendingOrders, order)
		}
	}

	return pendingOrders, nil
}

func (uc *OrderUseCase) CreateOrder(request *models.Order) (*models.Order, *model.ErrorLog) {
	result := make(chan models.OrderChan)

	go uc.orderRepository.Insert(uc.ctx, *request, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Order, nil
}

func (uc *OrderUseCase) UpdateOrder(request *models.Order) (*models.Order, *model.ErrorLog) {
	result := make(chan models.OrderChan)

	go uc.orderRepository.Update(uc.ctx, *request, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Order, nil
}

func (uc *OrderUseCase) GetAllByUserID(userID string) ([]*models.Order, *model.ErrorLog) {

	result := make(chan models.OrdersChan)

	go uc.orderRepository.GetAllByUserID(uc.ctx, userID, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Orders, nil
}

func (uc *OrderUseCase) GetAllByCourierID(courierID string) ([]*models.Order, *model.ErrorLog) {
	result := make(chan models.OrdersChan)

	go uc.orderRepository.GetAllByCourierID(uc.ctx, courierID, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Orders, nil
}

func (uc *OrderUseCase) GetByID(id string) (*models.Order, *model.ErrorLog) {
	result := make(chan models.OrderChan)

	go uc.orderRepository.GetByID(uc.ctx, id, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Order, nil
}

func (uc *OrderUseCase) GetAllByUsername(username string) ([]*models.Order, *model.ErrorLog) {
	result := make(chan models.OrdersChan)

	go uc.orderRepository.GetAllByUsername(uc.ctx, username, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Orders, nil
}

func (uc *OrderUseCase) GetAllByCourierUsername(username string) ([]*models.Order, *model.ErrorLog) {
	result := make(chan models.OrdersChan)

	go uc.orderRepository.GetAllByCourierUsername(uc.ctx, username, result)
	response := <-result

	if response.Error != nil {
		return nil, response.ErrorLog
	}

	return response.Orders, nil
}
