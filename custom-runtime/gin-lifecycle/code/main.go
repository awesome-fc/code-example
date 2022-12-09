package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"app": "Gin",
		})
	})
	router.POST("/initialize", func(c *gin.Context) {
		rid := c.GetHeader("x-fc-request-id")
		fmt.Println("FC Initialize Start RequestId: " + rid)
		//do your things
		fmt.Println("FC Initialize End RequestId: " + rid)
		c.String(200, "OK")
	})
	router.GET("/pre-freeze", func(c *gin.Context) {
		rid := c.GetHeader("x-fc-request-id")
		fmt.Println("FC PreFreeze Start RequestId: " + rid)
		//do your things
		fmt.Println("FC PreFreeze End RequestId: " + rid)
		c.String(200, "OK")
	})
	router.GET("/pre-stop", func(c *gin.Context) {
		rid := c.GetHeader("x-fc-request-id")
		fmt.Println("FC PreStop Start RequestId: " + rid)
		//do your things
		fmt.Println("FC PreStop End RequestId: " + rid)
		c.String(200, "OK")
	})
	router.Run(":9000")
}
