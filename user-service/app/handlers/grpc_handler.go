package handlers

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	pb "github.com/gonszalito/mini_project-online-courier/global-utils/protos"

	grpcclient "github.com/gonszalito/mini_project-online-courier/user-service/app/grpc-client"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func MainGrpcHandler(mongod mongodb.IMongoDB, ctx context.Context) {
	addr := fmt.Sprintf(":%s", os.Getenv("GRPC_SERVICE_PORT"))

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		logrus.Fatalln("Error when listen server address", err.Error())
	}

	grpcServer := grpc.NewServer()
	oauthService := grpcclient.InitGrpcOauthClient(mongod, ctx)

	pb.RegisterOauthServiceServer(grpcServer, oauthService)

	grpcServer.Serve(listen)

}
