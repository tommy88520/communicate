package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags Index
// @Success 200 {string} welcome
// @Router /user [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

func toChat(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(id)
	user.Identity = token

	c.JSON(200, gin.H{
		"code":    0,
		"message": "welcome",
		"data":    user,
	})
}
