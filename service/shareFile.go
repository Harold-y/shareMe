package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"shareMe/db"
	"shareMe/util"
)

func AddFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["fileList"]
	var prevNames []string
	shareCode := db.GenerateShareCode()
	if shareCode == "" {
		c.JSON(500, gin.H{"status": "failed"})
		return
	}
	for _, file := range files {
		err := c.SaveUploadedFile(file, util.UploadFolder+"/"+shareCode+"/"+file.Filename)
		if err != nil {
			c.JSON(500, gin.H{"status": "failed"})
			for _, name := range prevNames {
				err := os.Remove(util.UploadFolder + "/" + shareCode + "/" + name)
				if err != nil {
					fmt.Println("AddFile Delete Error!")
				}
			}
			return
		} else {
			prevNames = append(prevNames, file.Filename)
		}
	}

	res := db.CreateFileRecord(prevNames, shareCode)
	if !res {
		c.JSON(500, gin.H{"status": "failed"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "shareCode": shareCode})
}

func GetFile(c *gin.Context) {
	shareCode := c.Param("shareCode")
	if db.QueryShareCodeExist(shareCode) {
		files := db.QueryFileByShareCode(shareCode)
		c.JSON(200, gin.H{"status": "success", "files": files})
	} else {
		c.JSON(500, gin.H{"err": "Incorrect ShareCode"})
	}
}

func DownloadFile(c *gin.Context) {
	// shareCode := c.Param("shareCode")
	shareCode := c.Param("shareCode")
	fileName := c.Param("fileName")
	// Open
	_, errByOpenFile := os.Open(util.UploadFolder + "/" + shareCode + "/" + fileName)
	//非空处理
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "Cannot Open the File",
		})
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(util.UploadFolder + "/" + shareCode + "/" + fileName)
	return
}
