package routes

import (
	"user-athentication-golang/controllers"

	"github.com/gin-gonic/gin"
)

func ObjectRouter(app *gin.Engine) {
	app.POST("/object-detection", controllers.CreateObject())
	app.GET("/objects/:user_id", controllers.GetUserObjects())
	// app.GET("/objects/get/:item_id", controllers.GetUserObjects())
}
