package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/SouksavathPMS/go-basic-ecommerse/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

// NewApplication handles new application database connection
func NewApplication(
	productCollection *mongo.Collection,
	userCollection *mongo.Collection,
) *Application {
	return &Application{
		prodCollection: productCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Query("id")
		if productId == "" {
			log.Println("Product id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("Product id can not be empty"))
			return
		}

		userID := c.Query("user_id")
		if userID == "" {
			log.Println("User id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("User id can not be empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := database.AddProductToCart(
			ctx,
			app.prodCollection,
			app.userCollection,
			productID,
			userID,
		); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Product added!")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Query("id")
		if productId == "" {
			log.Println("Product id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("Product id can not be empty"))
			return
		}

		userID := c.Query("user_id")
		if userID == "" {
			log.Println("User id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("User id can not be empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := database.RemoveCartItem(
			ctx,
			app.prodCollection,
			app.userCollection,
			productID,
			userID,
		); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Product removed!")
	}
}

func (app *Application) GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Query("id")
		if productId == "" {
			log.Println("Product id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("Product id can not be empty"))
			return
		}

		userID := c.Query("user_id")
		if userID == "" {
			log.Println("User id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("User id can not be empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := database.InstantBuy(
			ctx,
			app.prodCollection,
			app.userCollection,
			productID,
			userID,
		); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Order placed successfully!")
	}
}
