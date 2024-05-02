package usecases

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/models"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignUp(request *models.User) (*models.User, *model.ErrorLog)
	Login(request *models.User) (*models.User, *model.ErrorLog)
	GetProfile(username string) (*models.User, *model.ErrorLog)
	Update(request *models.User) *model.ErrorLog
	Delete(request *models.User) *model.ErrorLog
	// Get(request *models.ProductRequest) (*models.Products, *model.ErrorLog)
	// GetDetail(uuid string) (*models.Product, *model.ErrorLog)
	// Insert(request *models.Product) (*models.Product, *model.ErrorLog)
	// Update(request *models.Product) *model.ErrorLog
	// Delete(request *models.Product) *model.ErrorLog
}

type UserUseCase struct {
	UserRepository repositories.IUserRepository
	mongod         mongodb.IMongoDB
	ctx            context.Context
}

func NewUserUseCase(
	userRepository repositories.IUserRepository,
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IUserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		mongod:         mongod,
		ctx:            ctx,
	}
}

func (u *UserUseCase) SignUp(
	request *models.User,
) (*models.User, *model.ErrorLog) {

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		return nil, errorLog
	}
	request.Password = string(hashedPassword)

	signUpUserChan := make(chan *models.UserChan)
	go u.UserRepository.SignUp(request, u.ctx, signUpUserChan)
	signUpUserResult := <-signUpUserChan

	if signUpUserResult.Error != nil {
		return nil, signUpUserResult.ErrorLog
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       signUpUserResult.User.ID,
		"username": request.Username,
		"role":     request.Role,
		"email":    request.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		return nil, errorLog
	}

	signUpUserResult.User.Token = tokenString

	if signUpUserResult.Error != nil {
		return nil, signUpUserResult.ErrorLog
	}

	return signUpUserResult.User, nil
}

// UserUseCase
func (u *UserUseCase) Login(request *models.User) (*models.User, *model.ErrorLog) {

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}
	// Retrieve user from repository by username

	loginUserChan := make(chan *models.UserChan)
	go u.UserRepository.GetUserByUsername(request.Username, u.ctx, loginUserChan)
	loginUserResult := <-loginUserChan

	if loginUserResult.Error != nil {
		// Handle repository errors
		return nil, loginUserResult.ErrorLog
	}

	user := loginUserResult.User

	// user, errorLog := u.UserRepository.GetUserByUsername(request.Username, u.ctx)
	// if errorLog != nil {
	// 	return nil, errorLog
	// }

	// Compare provided password with stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		errorLog := &model.ErrorLog{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid credentials",
			Err:        err,
		}
		return nil, errorLog
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusInternalServerError, "Failed to generate JWT token")
		return nil, errorLog
	}

	user.Token = tokenString

	println("user logged in successfully")

	return user, nil
}

func (u *UserUseCase) GetProfile(username string) (*models.User, *model.ErrorLog) {

	loginUserChan := make(chan *models.UserChan)
	go u.UserRepository.GetUserByUsername(username, u.ctx, loginUserChan)
	loginUserResult := <-loginUserChan

	loginUserResult.User.Password = ""

	if loginUserResult.Error != nil {
		// Handle repository errors
		return nil, loginUserResult.ErrorLog
	}

	// Return user profile
	return loginUserResult.User, nil
}

func (u *UserUseCase) Update(request *models.User) *model.ErrorLog {

	updateUserChan := make(chan *models.UserChan)
	go u.UserRepository.UpdateByUsername(request, u.ctx, updateUserChan)
	updateUserResult := <-updateUserChan

	if updateUserResult.Error != nil {
		return updateUserResult.ErrorLog
	}

	return nil
}

func (u *UserUseCase) Delete(request *models.User) *model.ErrorLog {

	deleteUserChan := make(chan *models.UserChan)
	go u.UserRepository.DeleteByUsername(request.Username, u.ctx, deleteUserChan)
	deleteUserResult := <-deleteUserChan

	if deleteUserResult.Error != nil {
		return deleteUserResult.ErrorLog
	}

	return nil
}
