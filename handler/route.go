package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouteAPI(router *gin.RouterGroup) {
	route := router.Group("/route")
	{
		route.GET("list-to-station/:FromStationCode/:EndStationLineID", listToStation)
		route.GET("list/:FromStationCode/:ToStationCode", listAllRoute)
		route.GET("detail/:RouteID", routeDetail)
	}
}