package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/SouksavathPMS/go-basic-ecommerse/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddAdress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid code"})
			c.Abort()
			return
		}

		address, err := bson.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
		}

		var addresses models.Address
		addresses.AddressID = primitive.NewObjectID()

		if err = c.BindJSON(addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		match_filer := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filer, unwind, group})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
		}
		var addressInfo []bson.M
		if err = pointCursor.All(ctx, &addressInfo); err != nil {
			panic(err)
		}

		var size int32

		for _, address_no := range addressInfo {
			count := address_no["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{bson.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{{Key: "address", Value: addresses}}}}
		} else {
			c.IndentedJSON(http.StatusBadRequest, "Not allowed")
		}
		ctx.Done()
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid search index"})
			c.Abort()
			return
		}
		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filteredUser := bson.D{bson.E{Key: "_id", Value: user_id}}
		update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filteredUser, update)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, "Wrong command")
			return
		}
		ctx.Done()
		c.IndentedJSON(http.StatusOK, "Successfully Deleted")
	}
}
