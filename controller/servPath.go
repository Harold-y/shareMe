package controller

import (
	"github.com/gin-gonic/gin"
	"shareMe/service"
	"shareMe/util"
)

func SetPath(router *gin.Engine) {
	// Link Redirect
	linkRoute := router.Group("/u")
	linkRoute.POST("/u/add", service.AddLink)
	linkRoute.GET("/u/:accessCode", service.GetLink)

	// Access File
	fileRoute := router.Group("/")
	fileRoute.POST("/add", service.AddFile)
	fileRoute.GET("/:accessCode", service.GetFile)

	// Destroy
	dbRoute := router.Group("/admin")
	dbRoute.POST("/destroy", util.Destroy)
}
