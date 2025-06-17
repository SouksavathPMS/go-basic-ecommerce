package routes

import (
	"github.com/SouksavathPMS/go-basic-ecommerse/controllers"
	"github.com/gin-gonic/gin"
)

// UserRoutes handles for /users routes
func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/users/signup", controllers.Signup())
	incommingRoutes.POST("/users/login", controllers.Login())
	incommingRoutes.GET("/users/products", controllers.SearchProduct())
	incommingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incommingRoutes.POST("/admin/add-product", controllers.ProductViewerAdmin())
}
