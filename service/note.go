package service

import (
	"github.com/gin-gonic/gin"
	"shareMe/db"
)

func AddNote(c *gin.Context) {
	shareCode := db.GenerateShareCode()
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil || param["Note"] == nil || param["Note"] == "" {
		c.JSON(500, gin.H{"err": err})
		return
	}
	if shareCode == "" {
		c.JSON(500, gin.H{"status": "failed"})
		return
	}
	res := db.CreateOtherRecord(shareCode, param["Note"].(string))
	if !res {
		c.JSON(500, gin.H{"err": err})
		return
	}
	c.JSON(200, gin.H{"status": "success", "shareCode": shareCode})
}

func GetNote(c *gin.Context) {
	shareCode := c.Param("shareCode")
	if db.QueryShareCodeExist(shareCode) {
		note := db.QueryOtherByShareCode(shareCode)
		c.JSON(200, gin.H{"status": "success", "note": note})
	} else {
		c.JSON(500, gin.H{"err": "Incorrect ShareCode"})
	}
}
