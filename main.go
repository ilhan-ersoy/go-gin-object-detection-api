package main

import (
	routes "user-athentication-golang/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {

	app := gin.New()
	app.Use(gin.Logger())

	routes.AuthRoutes(app)
	routes.UserRoutes(app)
	routes.ObjectRouter(app)

	// API-2
	app.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// API-1
	app.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	err := app.Run(":8080")

	if err != nil {
		return
	}
}
