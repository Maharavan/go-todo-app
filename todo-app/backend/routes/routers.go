package routes

import (
	"fmt"
	"todoapp/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eng *gin.Engine) {
	fmt.Println("POST newly created task")
	eng.POST("/postnewtask", controllers.PostnewTask)
	eng.GET("/gettask", controllers.GetTaskfromDB)
	eng.DELETE("/clearalltask", controllers.DeleteTaskfromDB)
	eng.PUT("/updatetask", controllers.UpdateTaskintoDB)
	eng.DELETE("/deletetask",controllers.DeleteTaskIDfromDB)
}
