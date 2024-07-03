package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
	"myapp/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/api")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You have access to this resource"})
		})

		userResponses := authorized.Group("user-responses")
		{
			userResponses.POST("/:user_id", controllers.SaveUserResponses)
			userResponses.PUT("/:user_id", controllers.UpdateUserResponses)
			userResponses.GET("/:user_id", controllers.GetUserResponses)
		}
		authorized.GET("/recommendations/:user_id", controllers.GetRecommendations)
		authorized.POST("/wallet", middlewares.AuthMiddleware(), controllers.AddToWallet)
		authorized.GET("/wallet/:user_id", middlewares.AuthMiddleware(), controllers.GetWallet)
	}
}
