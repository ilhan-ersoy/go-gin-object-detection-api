package main

import (
	routes "user-athentication-golang/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {

	app := gin.New()
	app.Use(gin.Logger())
	routes.ObjectRouter(app)
	routes.AuthRoutes(app)
	routes.UserRoutes(app)

	err := app.Run(":8080")

	if err != nil {
		return
	}
}
