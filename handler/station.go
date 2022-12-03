package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupStationAPI(router *gin.RouterGroup) {
	router.GET("/station/:LineId", listStation)
	router.GET("/station/:LineId/:StationCode", stationDetail)
}