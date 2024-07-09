package routes

import (
	"stock_prediction_backend/controllers"
	"stock_prediction_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/test", controllers.TestEndpoint) // Aggiungi questo endpoint di test

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
		authorized.GET("/profile/:user_id", controllers.GetProfile)
		authorized.GET("/recommendations/:user_id", controllers.GetRecommendations)
		authorized.POST("/wallet/:user_id", controllers.AddToWallet)
		authorized.GET("/wallet/:user_id", controllers.GetWallet)
		authorized.POST("/create-portfolio", controllers.CreatePortfolio)
	}
}
