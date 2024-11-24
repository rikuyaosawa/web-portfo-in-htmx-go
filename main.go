package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.Static("/static", "./static")

	server.GET("/", getLayoutTemplate)

	server.Run(":1010")
}

func getLayoutTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.tmpl", gin.H{
		"title": "Web Portfolio in HTMX/Go",
	})
}