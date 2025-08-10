package main

import (
	"fmt"
	"todoapp/renderer"
	"todoapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to TodoApp")
	r := gin.Default()

	renderer.Import_HTML(r)
	routes.RegisterRoutes(r)
	r.Run(":3001")
}
