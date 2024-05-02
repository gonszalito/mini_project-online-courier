package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"

	baseModel "github.com/gonszalito/mini_project-online-courier/global-utils/model"
	pb "github.com/gonszalito/mini_project-online-courier/global-utils/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func OauthMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {

		grpcAuthHost := fmt.Sprintf("%s:%s", os.Getenv("USER_SERVICE_HOST"), os.Getenv("USER_SERVICE_PORT"))

		var grpcConnection *grpc.ClientConn

		conn, err := grpc.DialContext(c, grpcAuthHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {

			errLog := &baseModel.ErrorLog{
				SystemMessage: err.Error(),
				Message:       "Grpc Initial Dial failed",
			}

			response := &baseModel.Response{
				Error:      errLog,
				StatusCode: http.StatusInternalServerError,
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		grpcConnection = conn

		authorizationBearerToken := helper.GetAuthorizationValue(c.GetHeader("Authorization"))

		if authorizationBearerToken == "" {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		ctx, _ := context.WithTimeout(c, 10*time.Second)
		client := pb.NewOauthServiceClient(grpcConnection)
		verifyAccessToken, err := client.ValidateToken(ctx, &pb.ValidateTokenRequest{Token: authorizationBearerToken})

		if err != nil {

			errLog := &baseModel.ErrorLog{
				SystemMessage: err.Error(),
				Message:       "Grpc Dial failed Call Service",
			}

			response := &baseModel.Response{
				Error:      errLog,
				StatusCode: http.StatusInternalServerError,
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		if verifyAccessToken.StatusCode == 200 {

			c.Writer.Header().Set("Content-Type", "application/json")

			user := verifyAccessToken.GetData()

			jwtTokenData := &baseModel.JWTAccessTokenPayload{
				Id:       user.Id,
				Role:     user.Role,
				Email:    user.Email,
				Username: user.Username,
				Name:     user.Name,
				Token:    authorizationBearerToken,
			}

			c.Set("user", jwtTokenData)

			c.Next()
		} else {

			errLog := &baseModel.ErrorLog{
				SystemMessage: verifyAccessToken.Error.SystemMessage,
				Message:       verifyAccessToken.Error.Message,
			}

			response := &baseModel.Response{
				Error:      errLog,
				StatusCode: int(verifyAccessToken.StatusCode),
			}

			c.AbortWithStatusJSON(int(verifyAccessToken.StatusCode), response)
			return
		}
	}
}
