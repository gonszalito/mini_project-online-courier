package repositories

import (
	"context"
	"net/http"
	"time"

	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	baseModel "github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderRepository interface {
	Insert(ctx context.Context, order models.Order, result chan models.OrderChan)
	Update(ctx context.Context, order models.Order, result chan models.OrderChan)
	GetAll(ctx context.Context, result chan models.OrdersChan)
	GetByID(ctx context.Context, id string, result chan models.OrderChan)
	GetAllByUserID(ctx context.Context, userID string, result chan models.OrdersChan)
	GetAllByCourierID(ctx context.Context, courierID string, result chan models.OrdersChan)
	GetAllByUsername(ctx context.Context, username string, result chan models.OrdersChan)
	GetAllByCourierUsername(ctx context.Context, username string, result chan models.OrdersChan)
}

type OrderRepository struct {
	mongodb mongodb.IMongoDB
	// logger         log.Logger
	collectionOrder string
}

func NewOrderRepository(mongodb mongodb.IMongoDB) IOrderRepository {
	return &OrderRepository{
		mongodb:         mongodb,
		collectionOrder: baseModel.COLLECTION_ORDER,
	}
}

func (r *OrderRepository) Insert(ctx context.Context, order models.Order, result chan models.OrderChan) {
	response := models.OrderChan{}

	now := time.Now()

	order.StartTime = &now
	order.Status = "pending"

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	insertResult, err := collection.InsertOne(ctx, order)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Order = &order
	response.Order.ID = insertResult.InsertedID.(primitive.ObjectID)
	result <- response
}

func (r *OrderRepository) Update(ctx context.Context, order models.Order, result chan models.OrderChan) {
	response := models.OrderChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	filter := bson.M{"_id": order.ID}

	timeNow := time.Now()

	update := bson.M{
		"status": order.Status,
	}

	if len(order.CourierID) > 0 {
		update = bson.M{
			"courier_id": order.CourierID,
			"status":     order.Status,
		}

	}

	if order.Status == "finished" || order.Status == "cancelled" {
		order.EndTime = &timeNow

		if len(order.CourierID) > 0 {
			update = bson.M{
				"courier_id": order.CourierID,
				"status":     order.Status,
				"end_time":   order.EndTime,
			}
		} else {
			update = bson.M{
				"status":   order.Status,
				"end_time": order.EndTime,
			}
		}
	}

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update}) // Pass the update variable with $set operator

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Order = &order

	// Set the ID of the modified document
	response.Order.ID = order.ID

	result <- response
}

func (r *OrderRepository) GetAll(ctx context.Context, result chan models.OrdersChan) {
	response := models.OrdersChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	var orders []*models.Order

	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			response.Error = err
			response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			result <- response
			return
		}
		orders = append(orders, &order)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Orders = orders
	result <- response
}

func (r *OrderRepository) GetByID(ctx context.Context, id string, result chan models.OrderChan) {
	response := models.OrderChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&response.Order)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	result <- response
}

func (r *OrderRepository) GetAllByUserID(ctx context.Context, userID string, result chan models.OrdersChan) {
	response := models.OrdersChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	filter := bson.M{"user_id": userID}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}
	defer cursor.Close(ctx)

	var orders []*models.Order

	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			response.Error = err
			response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			result <- response
			return
		}
		orders = append(orders, &order)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Orders = orders
	response.Error = nil
	result <- response
}

func (r *OrderRepository) GetAllByCourierID(ctx context.Context, courierID string, result chan models.OrdersChan) {
	response := models.OrdersChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	filter := bson.M{"courier_id": courierID}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	defer cursor.Close(ctx)

	var orders []*models.Order

	// Iterate over the result set and decode each document into the OrderWithID struct
	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			response.Error = err
			response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			result <- response
			return
		}
		orders = append(orders, &order)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Orders = orders
	result <- response
}

func (r *OrderRepository) GetAllByUsername(ctx context.Context, username string, result chan models.OrdersChan) {
	response := models.OrdersChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	filter := bson.M{"user_id": username}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	var orders []*models.Order

	err = cursor.All(ctx, &orders)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Orders = orders
	result <- response
}

func (r *OrderRepository) GetAllByCourierUsername(ctx context.Context, username string, result chan models.OrdersChan) {
	response := models.OrdersChan{}

	collection := r.mongodb.Client().Database(baseModel.DATABASE).Collection(r.collectionOrder)

	filter := bson.M{"courier_username": username}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	var orders []*models.Order

	err = cursor.All(ctx, &orders)

	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.Orders = orders
	result <- response
}
