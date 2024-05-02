package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/getsentry/sentry-go"
	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/order-service/app/handlers"
	envConfig "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	arg := os.Args[0]

	switch arg {
	case "main":
		mainWithoutArg()
		break
	default:
		mainWithoutArg()
		break
	}
}

func mainWithoutArg() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := envConfig.Load(".env"); err != nil {
		errStr := fmt.Sprintf(".env not load properly %s", err.Error())
		helper.SetSentryError(err, errStr, sentry.LevelError)
		panic(err)
	}

	ctx := context.Background()

	// mongoDB
	mongoDb := mongodb.NewMongoDB(mongodb.MongoDBParam{
		// Local: true,
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
	})

	defer mongoDb.Client().Disconnect(ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Printf("Starting Product Service HTTP Handler\n")
		handlers.MainHTTPHandler(mongoDb, ctx)
	}()

	wg.Wait()
}
