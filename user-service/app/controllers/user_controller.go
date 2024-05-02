package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/models"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/repositories"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/usecases"
)

type IUserController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type UserController struct {
	ctx              context.Context
	mongod           mongodb.IMongoDB
	userUseCase      usecases.IUserUseCase
	userRepositories repositories.IUserRepository
}

func NewUserController(
	ctx context.Context,
	mongod mongodb.IMongoDB,
	userUseCase usecases.IUserUseCase,
	userRepositories repositories.IUserRepository,
) IUserController {
	return &UserController{
		ctx:              ctx,
		mongod:           mongod,
		userUseCase:      userUseCase,
		userRepositories: userRepositories,
	}
}

func (c *UserController) SignUp(ctx *gin.Context) {
	var result model.Response
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	userResponse, errorLog := c.userUseCase.SignUp(&user)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = userResponse
	result.StatusCode = http.StatusCreated
	result.Message = "User created successfully"

	ctx.Header("Authorization", "Bearer "+userResponse.Token)
	ctx.JSON(http.StatusCreated, result)
}

// UserController
func (c *UserController) Login(ctx *gin.Context) {
	var result model.Response
	var user models.User

	// Bind JSON request body to user struct
	if err := ctx.BindJSON(&user); err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	// Call user use case to authenticate user
	userResponse, errorLog := c.userUseCase.Login(&user)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Message = "User authenticated successfully"

	// Set JWT token in response header
	ctx.Header("Authorization", "Bearer "+userResponse.Token)

	// Return success response with token
	result.Data = userResponse
	result.StatusCode = http.StatusOK
	ctx.JSON(http.StatusOK, result)
}

func (c *UserController) GetProfile(ctx *gin.Context) {

	var result model.Response

	// Get user ID from URL parameter
	userID := ctx.Param("username")

	userProfile, errorLog := c.userUseCase.GetProfile(userID)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = userProfile
	result.StatusCode = http.StatusOK
	// Return user profile as JSON response
	ctx.JSON(http.StatusOK, result)
}

func (c *UserController) Update(ctx *gin.Context) {
	var result model.Response
	var user models.User

	// Bind JSON request body to user struct
	if err := ctx.BindJSON(&user); err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	// Call use case to update user profile
	errorLog := c.userUseCase.Update(&user)
	if errorLog != nil {
		// Handle error response
		ctx.JSON(errorLog.StatusCode, errorLog)
		return
	}

	// Return success response
	result.Message = "User profile updated successfully"
	ctx.JSON(http.StatusOK, result)
}

func (c *UserController) Delete(ctx *gin.Context) {
	var result model.Response
	var user models.User

	// Bind JSON request body to user struct
	if err := ctx.BindJSON(&user); err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	// Call use case to delete user
	errorLog := c.userUseCase.Delete(&user)
	if errorLog != nil {
		// Handle error response
		ctx.JSON(errorLog.StatusCode, errorLog)
		return
	}

	// Return success response
	result.Message = "User deleted successfully"
	ctx.JSON(http.StatusOK, result)

}
