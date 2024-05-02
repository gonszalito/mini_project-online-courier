package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDD struct {
	client *mongo.Client
}

type IMongoDB interface {
	Client() *mongo.Client
}

type MongoDBParam struct {
	Host     string
	Port     int
	User     string
	Password string
	Local    bool
}

func NewMongoDB(param MongoDBParam) IMongoDB {
	// var mongoURL string

	connectionURI := fmt.Sprintf("mongodb+srv://jonathan:%s@user-service.2wljwcx.mongodb.net/?retryWrites=true&w=majority&appName=user-service", param.Password)

	clientOptions := options.Client().ApplyURI(connectionURI)

	// if param.Local {
	// 	mongoURL = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	// } else {
	// 	mongoURL = fmt.Sprintf("mongodb+srv://%s:%s@%s", param.User, param.Password, param.Host)

	// }

	client, err := mongo.Connect(context.Background(), clientOptions)

	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))

	if err != nil {

		panic(err)
	}

	return &MongoDD{
		client: client,
	}
}

func (m *MongoDD) Client() *mongo.Client {
	return m.client
}
