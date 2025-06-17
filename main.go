package main

import (
	"log"
	"os"

	"github.com/SouksavathPMS/go-basic-ecommerse/controllers"
	"github.com/SouksavathPMS/go-basic-ecommerse/database"
	"github.com/SouksavathPMS/go-basic-ecommerse/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// 1. **CREATE THE COLLECTION INSTANCES HERE**
	//    database.Client is already initialized in database.init()

	// 2. Pass these *created instances* to the NewApplication constructor
	app := controllers.NewApplication(
		database.ProdCollection,
		database.UserCollection,
	)

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)

	// router.Use(middleware.Authentication())
	router.GET("/add-to-cart", app.AddToCart())
	router.GET("/remove-item", app.RemoveItem())
	router.GET("/cart-checkout", app.BuyFromCart())
	router.GET("/instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
