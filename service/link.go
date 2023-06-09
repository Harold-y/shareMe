package service

import "github.com/gin-gonic/gin"

func AddLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetLink(c *gin.Context) {

}
