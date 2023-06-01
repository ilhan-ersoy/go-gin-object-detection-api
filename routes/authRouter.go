package routes

import (
	controller "go-gin-object-detection-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(app *gin.Engine) {
	app.POST("/users/signup", controller.SignUp())
	app.POST("/users/login", controller.Login())

}
