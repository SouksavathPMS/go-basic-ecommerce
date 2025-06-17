package controllers

import (
	"github.com/gin-gonic/gin"
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

func (app *Application) AddToCart() gin.HandlerFunc

func (app *Application) RemoveItem() gin.HandlerFunc

func (app *Application) GetItemFromCart() gin.HandlerFunc

func (app *Application) BuyFromCart() gin.HandlerFunc

func (app *Application) InstantBuy() gin.HandlerFunc
