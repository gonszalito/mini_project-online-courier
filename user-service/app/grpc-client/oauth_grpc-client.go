package grpcclient

import (
	"context"

	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	pb "github.com/gonszalito/mini_project-online-courier/global-utils/protos"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/usecases"
)

type OAuthGrpcClient struct {
	pb.UnimplementedOauthServiceServer
	oauthUseCase usecases.OAuthUseCaseInterface
}

func InitOAuthGrpcClient(oauthUseCase usecases.OAuthUseCaseInterface) *OAuthGrpcClient {
	return &OAuthGrpcClient{oauthUseCase: oauthUseCase}

}

func (c *OAuthGrpcClient) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.Response, error) {
	var result = &pb.Response{}

	validateToken, errLog := c.oauthUseCase.ValidateToken(ctx, request.Token)

	if errLog != nil {

		errGrpc := &pb.ErrorResponse{
			Message:       errLog.Message.(string),
			SystemMessage: errLog.SystemMessage,
			StatusCode:    int32(errLog.StatusCode),
		}

		result.Error = errGrpc
		result.StatusCode = int32(errLog.StatusCode)
		result.Data = nil
		return result, nil
	}

	var validateTokenGrpc pb.ValidateTokenResponse
	_ = helper.DecodeMapType(validateToken, &validateTokenGrpc)

	result.Data = &validateTokenGrpc
	result.StatusCode = 200
	result.Error = nil

	return result, nil

}
