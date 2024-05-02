package repositories

import (
	"context"
	"net/http"

	"github.com/gonszalito/mini_project-online-courier/user-service/app/constants"
	"github.com/gonszalito/mini_project-online-courier/user-service/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gonszalito/mini_project-online-courier/global-utils/helper"
	"github.com/gonszalito/mini_project-online-courier/global-utils/model"
	"github.com/gonszalito/mini_project-online-courier/global-utils/mongodb"
)

type IUserRepository interface {
	SignUp(user *models.User, ctx context.Context, result chan *models.UserChan) // Sign up new user
	GetUserByUsername(username string, ctx context.Context, result chan *models.UserChan)
	DeleteByUsername(username string, ctx context.Context, result chan *models.UserChan)
	UpdateByUsername(user *models.User, ctx context.Context, result chan *models.UserChan)
	// Get(request *models.ProductRequest, ctx context.Context, result chan *models.ProductsChan) // Get all products
	// GetDetail(uuid string, ctx context.Context, result chan *models.ProductChan)               // Get product detail
	// Insert(product *models.Product, ctx context.Context, result chan *models.ProductChan)      // Insert new product
	// Update(product *models.Product, ctx context.Context, result chan *models.ProductChan)      // Update product
	// Delete(product *models.Product, ctx context.Context, result chan *models.ProductChan)      // Delete product
}

type UserRepository struct {
	mongodb mongodb.IMongoDB
	// logger         log.Logger
	collectionUser string
}

func NewUserRepository(
	mongodb mongodb.IMongoDB,
) IUserRepository {
	return &UserRepository{
		mongodb:        mongodb,
		collectionUser: constants.COLLECTION_USER,
	}
}

func (r *UserRepository) SignUp(
	user *models.User,
	ctx context.Context,
	result chan *models.UserChan,
) {

	// jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	// if len(jwtKey) == 0 {
	// 	log.Fatal("JWT_SECRET_KEY environment variable not set")
	// }

	response := &models.UserChan{}

	collection := r.mongodb.Client().Database(constants.DATABASE).Collection(r.collectionUser)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": user.Username,
	// 	"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	// })

	// tokenString, err := token.SignedString(jwtKey)
	// if err != nil {
	// 	errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
	// 	response.Error = err
	// 	response.ErrorLog = errorLog
	// 	result <- response
	// }

	insertedUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Role:     user.Role,
		// Token:    tokenString,
	}

	insertResult, err := collection.InsertOne(ctx, insertedUser)
	if err != nil {
		response.Error = err
		response.ErrorLog = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		result <- response
		return
	}

	response.User = &insertedUser

	response.User.ID = insertResult.InsertedID.(primitive.ObjectID)
	result <- response
	// return
}

// UserRepository
// UserRepository
func (r *UserRepository) GetUserByUsername(username string, ctx context.Context, result chan *models.UserChan) {

	collection := r.mongodb.Client().Database(constants.DATABASE).Collection(r.collectionUser)

	var user models.User
	filter := bson.M{"username": username}

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			result <- &models.UserChan{
				User:  nil,
				Error: err,
				ErrorLog: &model.ErrorLog{
					StatusCode: http.StatusNotFound,
					Message:    "User not found",
					Err:        err,
				},
			}
			return
		}
		result <- &models.UserChan{
			User:  nil,
			Error: err,
			ErrorLog: &model.ErrorLog{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to retrieve user",
				Err:        err,
			},
		}

		return
	}

	// Send the retrieved user via the result channel
	result <- &models.UserChan{
		User:     &user,
		Error:    nil,
		ErrorLog: nil,
	}
}

func (r *UserRepository) DeleteByUsername(username string, ctx context.Context, result chan *models.UserChan) {
	defer close(result) // Ensure the result channel is closed when this function exits

	collection := r.mongodb.Client().Database(constants.DATABASE).Collection(r.collectionUser)

	filter := bson.M{"username": username}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		result <- &models.UserChan{
			User:  nil,
			Error: err,
			ErrorLog: &model.ErrorLog{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to delete user",
				Err:        err,
			},
		}
		return
	}

	result <- &models.UserChan{
		User:     nil,
		Error:    nil,
		ErrorLog: nil,
	}
}

func (r *UserRepository) UpdateByUsername(user *models.User, ctx context.Context, result chan *models.UserChan) {
	defer close(result) // Ensure the result channel is closed when this function exits

	collection := r.mongodb.Client().Database(constants.DATABASE).Collection(r.collectionUser)

	filter := bson.M{"username": user.Username}
	update := bson.M{"$set": bson.M{
		"email":    user.Email,
		"password": user.Password,
		"name":     user.Name,
		"role":     user.Role,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		result <- &models.UserChan{
			User:  nil,
			Error: err,
			ErrorLog: &model.ErrorLog{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to update user",
				Err:        err,
			},
		}
		return
	}

	result <- &models.UserChan{
		User:     nil,
		Error:    nil,
		ErrorLog: nil,
	}
}
