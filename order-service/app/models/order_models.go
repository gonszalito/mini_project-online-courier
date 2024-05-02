package models

import (
	"time"

	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID        string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CourierID     string             `json:"courier_id,omitempty" bson:"courier_id,omitempty"`
	Status        string             `json:"status,omitempty" bson:"status,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Price         int64              `json:"price,omitempty" bson:"price,omitempty"`
	StartLocation string             `json:"start_location,omitempty" bson:"start_location,omitempty"`
	EndLocation   string             `json:"end_location,omitempty" bson:"end_location,omitempty"`
	StartTime     *time.Time         `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime       *time.Time         `json:"end_time,omitempty" bson:"end_time,omitempty"`
}

type OrderChan struct {
	Order      *Order          `json:"order,omitempty"`
	Total      int64           `json:"total,omitempty"`
	Error      error           `json:"error,omitempty"`
	ErrorLog   *model.ErrorLog `json:"error_log,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}

type OrdersChan struct {
	Orders     []*Order        `json:"order,omitempty"`
	Total      int64           `json:"total,omitempty"`
	Error      error           `json:"error,omitempty"`
	ErrorLog   *model.ErrorLog `json:"error_log,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}
