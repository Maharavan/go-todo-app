package routes

import (
	"fmt"
	"todoapp/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eng *gin.Engine) {
	fmt.Println("POST newly created task")
	eng.POST("/updatetask", controllers.PostnewTask)
	eng.GET("/gettask", controllers.GetTaskfromDB)

}
