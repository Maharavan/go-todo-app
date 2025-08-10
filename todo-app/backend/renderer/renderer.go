package renderer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Import_HTML(eng *gin.Engine) {
	fmt.Println("Importing Html Files")
	eng.LoadHTMLGlob("../frontend/*.html")
	eng.StaticFile("/script.js", "../frontend/script.js")
	eng.StaticFile("/styles.css", "../frontend/styles.css")
	eng.Static("/assets", "../frontend/assets/")
	eng.GET("/", html_parser)

}

func html_parser(router *gin.Context) {
	router.HTML(http.StatusOK, "index.html", gin.H{
		"title": "To-Do-App",
	})
}
