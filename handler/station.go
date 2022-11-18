package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupStationAPI(router *gin.RouterGroup) {
	router.GET("/station", listStation)
	router.GET("/station/:stationId", stationDetail)
}