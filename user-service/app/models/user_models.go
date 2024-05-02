package models

import (
	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" `
	Name     string             `json:"name,omitempty" bson:"name,omitempty"  `
	Token    string             `json:"token,omitempty" `
	Role     string             `json:"role,omitempty"`
}

type UserChan struct {
	User       *User           `json:"user,omitempty"`
	Id         string          `json:"id,omitempty"`
	Error      error           `json:"error,omitempty"`
	ErrorLog   *model.ErrorLog `json:"error_log,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}
