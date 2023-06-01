package routes

import (
	controller "user-athentication-golang/controllers"
	"user-athentication-golang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	app.Use(middleware.Authentication())
	app.GET("/users", controller.GetUsers())
	app.GET("/users/:user_id", controller.GetUser())
}
