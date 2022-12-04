package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupTripAPI(router *gin.RouterGroup) {
	trip := router.Group("/trip")
	{
		trip.GET("/place/:RouteID", listPlace)
		trip.POST("/", createTrip)
		trip.GET("/:TripID", tripDetail)
		trip.GET("/", listTrip)
		// auth.PUT("/:TripID", updateTrip)
		// auth.DELETE("/:TripID", deleteTrip)
	}
}