package handler

import (
	"discord-bot/lib/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJoinUserData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var join []*db.UserJoin
		connection.Model(join).Order("created_at asc").Find(&join)

		c.JSON(http.StatusOK, join)

	}
}
