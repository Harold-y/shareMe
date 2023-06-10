package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"shareMe/db"
	"shareMe/util"
	"time"
)

func InnerDestroy() {
	// DELETE ALL RELEVANT FILE OR RECORDS
	currentTime := time.Now()

	fmt.Println("[", currentTime.Format("2006.01.02 15:04:05"), "]", "Begin to Clear the Memory, I'm Sorry, The Machine.")
	// Clear the Database
	iter := db.DB.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		db.PoolByShareCode(string(key[:]))
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println("Error Occurred during Destroy Iter Release!")
	}

	// Clear the disk folder
	err = os.RemoveAll(util.UploadFolder)
	if err != nil {
		fmt.Println("Error Occurred during Destroy folder delete!")
	}
	err = os.MkdirAll(util.UploadFolder, 0777)
	if err != nil {
		fmt.Println("Error Occurred during Destroy folder create!")
	}
}

func Destroy(c *gin.Context) {
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err != nil || param["adminToken"] == nil {
		c.JSON(500, gin.H{"err": err})
		return
	}
	// adminToken := c.PostForm("adminToken")
	if util.AdminToken == param["adminToken"] {
		InnerDestroy()
		c.JSON(200, gin.H{"status": "success"})
	} else {
		c.JSON(500, gin.H{"err": "wrong token"})
	}
}
