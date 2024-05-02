package usecases

import (
	"context"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	baseModel "github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/models"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/repositories"
)

type OAuthUseCaseInterface interface {
	ValidateToken(ctx context.Context, token string) (*baseModel.JWTAccessTokenPayload, *baseModel.ErrorLog)
	// GenerateToken(ctx context.Context, request *models.SignUpRequest) (string, *baseModel.ErrorLog)
}

type oauthUseCase struct {
	userRepository repositories.IUserRepository
}

func InitOAuthUseCase(userRepository repositories.IUserRepository) OAuthUseCaseInterface {
	return &oauthUseCase{
		userRepository: userRepository,
	}
}
func (u *oauthUseCase) ValidateToken(ctx context.Context, tokenString string) (*baseModel.JWTAccessTokenPayload, *baseModel.ErrorLog) {

	jwtKey := os.Getenv("JWT_SECRET_KEY")

	token, err := helper.VerifyToken(tokenString, jwtKey)

	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusUnauthorized, err.Error())
		return nil, errorLog
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		errorLog := helper.WriteLog(err, http.StatusUnauthorized, "failed to extract claims from token")
		return nil, errorLog
	}

	// Extract username from claims
	username, ok := claims["username"].(string)
	if !ok {
		errorLog := helper.WriteLog(err, http.StatusUnauthorized, "failed to extract username from token claims")
		return nil, errorLog
	}

	// Fetch user details from the repository based on the username
	userChan := make(chan *models.UserChan)
	go u.userRepository.GetUserByUsername(username, ctx, userChan)
	userResult := <-userChan

	// Check if there was an error fetching the user
	if userResult.Error != nil {
		return nil, userResult.ErrorLog
	}

	// Extract user information from the retrieved user
	user := userResult.User
	if user == nil {
		// User not found
		return nil, &baseModel.ErrorLog{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
		}
	}

	// Populate JWTAccessTokenPayload with user details
	jwtPayload := &baseModel.JWTAccessTokenPayload{
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		Id:       user.ID.Hex(),
		// You may populate other fields of the JWTAccessTokenPayload struct if needed
	}

	return jwtPayload, nil
}
