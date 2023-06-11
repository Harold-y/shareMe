package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shareMe/db"
)

func AddLink(c *gin.Context) {
	shareCode := db.GenerateShareCode()
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil || param["URL"] == nil || param["URL"] == "" {
		c.JSON(500, gin.H{"err": err})
		return
	}
	if shareCode == "" {
		c.JSON(500, gin.H{"status": "failed"})
		return
	}
	res := db.CreateOtherRecord(shareCode, param["URL"].(string), "URL")
	if !res {
		c.JSON(500, gin.H{"err": err})
		return
	}
	c.JSON(200, gin.H{"status": "success", "shareCode": shareCode})
}

func GetLink(c *gin.Context) {
	shareCode := c.Param("shareCode")
	if db.QueryShareCodeExist(shareCode) {
		link := db.QueryOtherByShareCode(shareCode)
		c.Redirect(http.StatusMovedPermanently, link)
	} else {
		c.JSON(500, gin.H{"err": "Incorrect ShareCode"})
	}
}
