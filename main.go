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
	server.GET("/about", getAboutTemplate)
	server.GET("/project", getProjectTemplate)
	server.GET("/blog", getBlogTemplate)

	server.Run(":1010")
}

func getLayoutTemplate(c *gin.Context) {

	contentData := gin.H{
		"title": "Web Portfolio in HTMX/Go",
		"header": "Welcome",
		"text": "Welcome to my portfolio built with HTMX and Go!",
	}

	c.HTML(http.StatusOK, "layout.tmpl", contentData)
}

func getAboutTemplate(c *gin.Context) {
	isHX := c.GetHeader("hx-request" )

	contentData := gin.H{
		"header": "About",
		"text": "About me",
	}

	if (isHX == "true") {
		c.HTML(http.StatusOK, "content.tmpl", contentData)
	} else {
		contentData["title"] = "Web Portfolio in HTMX/Go"
		c.HTML(http.StatusOK, "layout.tmpl", contentData)
	}
}

func getProjectTemplate(c *gin.Context) {
	isHX := c.GetHeader("hx-request" )

	contentData := gin.H{
		"header": "Project",
		"text": "Project list...",
	}

	if (isHX == "true") {
		c.HTML(http.StatusOK, "content.tmpl", contentData)
	} else {
		contentData["title"] = "Web Portfolio in HTMX/Go"
		c.HTML(http.StatusOK, "layout.tmpl", contentData)
	}
}

func getBlogTemplate(c *gin.Context) {
	isHX := c.GetHeader("hx-request" )

	contentData := gin.H{
		"header": "Blog",
		"text": "Blog posts...",
	}

	if (isHX == "true") {
		c.HTML(http.StatusOK, "content.tmpl", contentData)
	} else {
		contentData["title"] = "Web Portfolio in HTMX/Go"
		c.HTML(http.StatusOK, "layout.tmpl", contentData)
	}
}