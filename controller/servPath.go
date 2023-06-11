package controller

import (
	"github.com/gin-gonic/gin"
	"shareMe/service"
)

func SetPath(router *gin.Engine) {
	// Notes
	noteRoute := router.Group("/n")
	noteRoute.POST("/add", service.AddNote)
	noteRoute.GET("/:shareCode", service.GetNote)

	// Link Redirect
	linkRoute := router.Group("/u")
	linkRoute.POST("/add", service.AddLink)
	linkRoute.GET("/:shareCode", service.GetLink)

	// Access File
	fileRoute := router.Group("/")
	fileRoute.POST("/add", service.AddFile)
	fileRoute.GET("/:shareCode", service.GetFile)
	fileRoute.GET("/download/:shareCode/:fileName", service.DownloadFile)
	fileRoute.GET("/download/all/:shareCode", service.DownloadAllFile)

	// Destroy
	dbRoute := router.Group("/admin")
	dbRoute.POST("/destroy", service.Destroy)

	// Get All In One Place
	getRoute := router.Group("/get")
	getRoute.GET("/:shareCode", service.GetThing)
}
