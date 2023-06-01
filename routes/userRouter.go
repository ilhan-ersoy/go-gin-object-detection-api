package routes

import (
	controller "go-gin-object-detection-api/controllers"
	"go-gin-object-detection-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	app.Use(middleware.Authentication())
	app.GET("/users", controller.GetUsers())
	app.GET("/users/:user_id", controller.GetUser())
}
