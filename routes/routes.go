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

		authorized.POST("/questionnaire", controllers.SaveUserResponses)
		authorized.PUT("/questionnaire", controllers.UpdateUserResponses)
		authorized.GET("/questionnaire", controllers.GetUserResponses)
		//authorized.GET("/recommend", controllers.GetRecommendations)
		//authorized.POST("/recommend", controllers.Recommend)
		//authorized.POST("/predict", controllers.Predict)
	}
}
