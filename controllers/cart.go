package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/SouksavathPMS/go-basic-ecommerse/database"
	"github.com/SouksavathPMS/go-basic-ecommerse/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
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
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		user_id, _ := primitive.ObjectIDFromHex(userID)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledCart models.User
		err := UserCollection.FindOne(ctx, bson.D{bson.E{Key: "_id", Value: user_id}}).Decode(&filledCart)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "not found")
			return
		}

		filter_match := bson.D{{Key: "$match", Value: bson.D{bson.E{Key: "_id", Value: user_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user_cart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$user_cart.price"}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil {
			log.Println(err)
		}

		var listing []bson.M
		if err = pointCursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		for _, json := range listing {
			c.IndentedJSON(http.StatusOK, json["total"])
			c.IndentedJSON(http.StatusOK, filledCart.UserCart)

		}
		ctx.Done()

	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			log.Panicln("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.prodCollection, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Successfully placed order")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Query("id")
		if productId == "" {
			log.Panicln("Product id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("Product id can not be empty"))
			return
		}

		userID := c.Query("user_id")
		if userID == "" {
			log.Panicln("User id is empty")
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
