package main

import (
	"chat/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	go ws.Manager.Start()
	r := gin.Default()
	r.GET("/ws", ws.WsHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.LoadHTMLGlob("template/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "client.html", gin.H{})
	})

	r.Run(":8282")
}
