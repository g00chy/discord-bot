package handler

import (
	"discord-bot/lib/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAfkCallUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *[]db.User
		connection.Model(user).Order("created_at asc").Find(&user)

		c.JSON(http.StatusOK, user)
	}
}
