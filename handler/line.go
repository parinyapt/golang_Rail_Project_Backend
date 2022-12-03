package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupLineAPI(router *gin.RouterGroup) {
	router.GET("/line", listLine)
	// router.GET("/station/:stationId", stationDetail)
}