package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"shareMe/controller"
	"shareMe/db"
	"shareMe/service"
	"shareMe/util"
)

func main() {
	// Start the database
	db.StartDB("data")
	// Set up Config
	util.SetConfig()
	if util.DestroyOnStart == "Yes" {
		service.InnerDestroy()
	}
	// Print Start Information
	util.InformAtStart()
	// Set up Auto Destroyer
	go func() {
		err := gocron.Every(uint64(util.DailyDestroyHour)).Hours().Do(service.InnerDestroy)
		if err != nil {
			return
		}
		<-gocron.Start()
	}()

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.MaxMultipartMemory = int64(util.MaxUploadSize)
	controller.SetPath(r)
	err := r.Run(util.Port)
	if err != nil {
		return
	}
	<-gocron.Start()
}
