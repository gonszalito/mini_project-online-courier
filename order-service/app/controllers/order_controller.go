package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	baseModel "github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/models"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/repositories"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/usecases"
)

type IOrderController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllByUserID(ctx *gin.Context)
	GetAllByCourierID(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAllByUsername(ctx *gin.Context)
	GetAllByCourierUsername(ctx *gin.Context)
	GetAllPending(ctx *gin.Context)
}

type OrderController struct {
	ctx               context.Context
	mongod            mongodb.IMongoDB
	orderUseCase      usecases.IOrderUseCase
	orderRepositories repositories.IOrderRepository
}

func NewOrderController(
	ctx context.Context,
	mongod mongodb.IMongoDB,
	orderUseCase usecases.IOrderUseCase,
	orderRepositories repositories.IOrderRepository,
) IOrderController {
	return &OrderController{
		ctx:               ctx,
		mongod:            mongod,
		orderUseCase:      orderUseCase,
		orderRepositories: orderRepositories,
	}
}

func (ctrl *OrderController) GetAllPending(ctx *gin.Context) {
	var result model.Response

	orders, errorLog := ctrl.orderUseCase.GetAllPending()

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orders
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *OrderController) Create(ctx *gin.Context) {

	user := ctx.Value("user").(*baseModel.JWTAccessTokenPayload)

	var result model.Response
	var request models.Order

	request.UserID = user.Id

	err := ctx.BindJSON(&request)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	order, errorLog := ctrl.orderUseCase.CreateOrder(&request)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = order
	result.StatusCode = http.StatusCreated
	ctx.JSON(http.StatusCreated, result)
}

func (ctrl *OrderController) Update(ctx *gin.Context) {
	var result model.Response
	var request models.Order

	user := ctx.Value("user").(*baseModel.JWTAccessTokenPayload)

	if user.Role == "courier" {
		request.CourierID = user.Id
	} else {
		request.UserID = user.Id
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	order, errorLog := ctrl.orderUseCase.UpdateOrder(&request)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = order
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *OrderController) GetAllByUserID(ctx *gin.Context) {
	var result model.Response

	user := ctx.Value("user").(*baseModel.JWTAccessTokenPayload)

	if user.Role != "user" {
		errorLog := helper.WriteLog(helper.NewError("Unauthorized"), http.StatusUnauthorized, "Unauthorized")
		result.Data = errorLog
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	orders, errorLog := ctrl.orderUseCase.GetAllByUserID(user.Id)

	if errorLog != nil {

		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orders
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *OrderController) GetAllByCourierID(ctx *gin.Context) {
	var result model.Response

	user := ctx.Value("user").(*baseModel.JWTAccessTokenPayload)

	if user.Role != "courier" {
		println("eerror here")
		errorLog := helper.WriteLog(helper.NewError("Unauthorized"), http.StatusUnauthorized, "Unauthorized")
		result.Error = errorLog
		ctx.JSON(http.StatusUnauthorized, result)
		return
	}

	orders, errorLog := ctrl.orderUseCase.GetAllByCourierID(user.Id)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orders
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *OrderController) GetByID(ctx *gin.Context) {
	var result model.Response
	id := ctx.Param("id")

	order, errorLog := ctrl.orderUseCase.GetByID(id)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = order
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *OrderController) GetAllByUsername(ctx *gin.Context) {
	var result model.Response
	username := ctx.Param("username")

	orders, errorLog := ctrl.orderUseCase.GetAllByUsername(username)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orders
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (ctrl *OrderController) GetAllByCourierUsername(ctx *gin.Context) {
	var result model.Response
	username := ctx.Param("username")

	orders, errorLog := ctrl.orderUseCase.GetAllByCourierUsername(username)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orders
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}
