package handler

import (
	"discord-bot/lib/db"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJoinUserData() gin.HandlerFunc {
	return func(c *gin.Context) {

		var userJoin []*db.UserJoin
		connection.Model(userJoin).Order("created_at asc").Find(&userJoin)

		c.JSON(http.StatusOK, userJoin)

	}
}
