package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupStationAPI(router *gin.RouterGroup) {
	station := router.Group("/station")
	{
		station.GET("/:LineId", listStation)
		station.GET("/detail/:StationID", stationDetail)
	}
}