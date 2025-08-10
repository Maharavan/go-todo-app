package controllers

import (
	"fmt"
	"net/http"
	"todoapp/database"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func PostnewTask(rout *gin.Context) {
	fmt.Println("Post route")
	var T models.Todo
	if err := rout.ShouldBindJSON(&T); err != nil {
		fmt.Println(err)
		rout.JSON(http.StatusBadRequest, gin.H{
			"error": "Info not updated",
		})
		return
	}
	fmt.Printf("Task: %s, Status: %s", T.Task, T.Status)
	database.InsertTask(&T)
	rout.JSON(http.StatusOK, gin.H{"status": "Yes"})
}

func GetTaskfromDB(rout *gin.Context) {
	res, err := database.GetTask()
	if err != nil {
		rout.JSON(http.StatusNotAcceptable, gin.H{"tasks": nil})
	}
	rout.JSON(http.StatusOK, gin.H{"tasks": res})
}
