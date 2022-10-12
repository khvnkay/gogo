package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khvnkay/ecommerce-yt/controllers"
)

func UserRoutes(incomeRoutes *gin.Engine) {
	incomeRoutes.POST("/user/signup", controllers.Signup())
	incomeRoutes.POST("/user/login", controllers.Login())
	incomeRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomeRoutes.GET("/users/produxtview", controllers.SearchProduct())
	incomeRoutes.GET("/users/search", controllers.SearchProductByQuery())

}
