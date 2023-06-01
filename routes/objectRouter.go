package routes

import (
	"go-gin-object-detection-api/controllers"

	"github.com/gin-gonic/gin"
)

func ObjectRouter(app *gin.Engine) {
	app.POST("/object-detection", controllers.CreateObject())
	app.GET("/objects/:user_id", controllers.GetUserObjects())
	app.GET("/object/get/:item_id", controllers.GetObject())
	app.DELETE("/object/delete/:item_id", controllers.DeleteObject())
	app.DELETE("/objects/delete/:user_id", controllers.DeleteObjects())
}
