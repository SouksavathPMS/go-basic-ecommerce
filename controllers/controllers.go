package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SouksavathPMS/go-basic-ecommerse/database"
	"github.com/SouksavathPMS/go-basic-ecommerse/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	}
}

// Login handles user login
func Login() gin.HandlerFunc

// ProductViewerAdmin handles the product viewer for admin ony
func ProductViewerAdmin() gin.HandlerFunc

// SearchProduct handles product search
func SearchProduct() gin.HandlerFunc

// SearchProduct handles product search with its text query
func SearchProductByQuery() gin.HandlerFunc
