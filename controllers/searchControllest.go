package controllers

import (
	"e/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}
func Search(c *gin.Context) {
	c.JSON(200, gin.H{
		"movie": "TODO: Search",
	})
}
