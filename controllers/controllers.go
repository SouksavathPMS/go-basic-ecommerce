package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SouksavathPMS/go-basic-ecommerse/database"
	"github.com/SouksavathPMS/go-basic-ecommerse/models"
	"github.com/SouksavathPMS/go-basic-ecommerse/tokens"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var validate = validator.New()

// HashPassword handles password to be hash and return its hashed password
func HashPassword(password string) string {
	return ""
}

// VerifyPassword handles the user password comparision either correct or not with returning of the password
func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	return false, ""
}

// Signup handles user sign up
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		// Checking if the user with this email is already exist
		count, err := database.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User with this email already existed!",
			})
			return
		}

		// Checking if the user with this phone is already exist
		count, err = database.UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		if count > 0 {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User with this phone no. already existed!",
			})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.CreatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()

		// Generate and get both accesssToken and refreshToken
		token, refreshToken := tokens.TokenGenerater(user.UserID, *user.Email)
		user.Token = &token
		user.RefreshToken = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.AddressDetails = make([]models.Address, 0)
		user.OrderStatus = make([]models.Order, 0)

		_, insertErr := database.UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while trying to create user",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Create user successfully!",
		})
	}
}

// Login handles user login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
		err := database.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Invalid user credential",
			})
		}

		isValidPassword, _ := VerifyPassword(*user.Password, *foundUser.Password)
		if !isValidPassword {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid user credential",
			})
			return
		}
		// Generate and get both accesssToken and refreshToken
		token, refreshToken := tokens.TokenGenerater(foundUser.UserID, *foundUser.Email)

		// Update new user's token
		tokens.UpdateAllToken(token, refreshToken, foundUser.UserID)

		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successed!",
		})
	}
}

// ProductViewerAdmin handles the product viewer for admin ony
func ProductViewerAdmin() gin.HandlerFunc

// SearchProduct handles product search
func SearchProduct() gin.HandlerFunc

// SearchProduct handles product search with its text query
func SearchProductByQuery() gin.HandlerFunc
