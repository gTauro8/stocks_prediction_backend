package main

import (
	"github.com/gin-gonic/gin"
	"myapp/config"
	"myapp/routes"
)

func main() {
	config.Setup()

	r := gin.Default()
	routes.SetupRoutes(r)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}
