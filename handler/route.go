package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouteAPI(router *gin.RouterGroup) {
	route := router.Group("/route")
	{
		route.GET("list-to-station/:FromStationCode", listToStation)
		route.GET("all/:FromStationCode/:ToStationCode", listAllRoute)
		route.GET("detail/:FromStationCode/:ToStationCode", routeDetail)
	}
}