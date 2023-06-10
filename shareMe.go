package main

import (
	"github.com/gin-gonic/gin"
	"shareMe/controller"
	"shareMe/db"
	"shareMe/util"
)

func main() {
	// Start the database
	db.StartDB("data")
	// Set up Config
	util.SetConfig()

	r := gin.Default()
	r.MaxMultipartMemory = int64(util.MaxUploadSize)
	controller.SetPath(r)
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
