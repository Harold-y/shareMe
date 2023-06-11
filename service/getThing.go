package service

import (
	"github.com/gin-gonic/gin"
	"shareMe/db"
)

func GetThing(c *gin.Context) {
	shareCode := c.Param("shareCode")
	if db.QueryShareCodeExist(shareCode) {
		infoType := db.QueryShareCodeType(shareCode)
		if infoType == "File" {
			files := db.QueryFileByShareCode(shareCode)
			c.JSON(200, gin.H{"status": "success", "files": files})
		}
		if infoType == "URL" {
			link := db.QueryOtherByShareCode(shareCode)
			c.JSON(200, gin.H{"status": "success", "link": link})
		}
		if infoType == "Note" {
			note := db.QueryOtherByShareCode(shareCode)
			c.JSON(200, gin.H{"status": "success", "note": note})
		}
		if infoType == "Error" {
			c.JSON(500, gin.H{"err": "Incorrect ShareCode"})
			return
		}
	} else {
		c.JSON(500, gin.H{"err": "Incorrect ShareCode"})
	}
}
