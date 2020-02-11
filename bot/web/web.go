package main

import (
	"discord-bot/web/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user_join", handler.GetJoinUserData())
	r.GET("/user", handler.GetAllAfkCallUsers())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
