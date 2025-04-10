package routes

import (
	"online-store/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        users := api.Group("/users")
        {
            users.POST("/register", controllers.RegisterUser)
            users.POST("/login", controllers.LoginUser)
        }

        products := api.Group("/products")
        {
            products.POST("/", controllers.CreateProduct)
			products.GET("/", controllers.GetAllProducts)
			products.GET("/:id", controllers.GetProductByID)
			products.PUT("/:id", controllers.UpdateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
        }
    }

    return r
}
