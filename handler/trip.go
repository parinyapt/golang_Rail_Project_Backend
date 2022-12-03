package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupTripAPI(router *gin.RouterGroup) {
	auth := router.Group("/trip")
	{
		auth.GET("/place", listPlace)
		auth.POST("/", createTrip)
		auth.GET("/:TripID", tripDetail)
		auth.PUT("/:TripID", updateTrip)
		auth.DELETE("/:TripID", deleteTrip)
	}
}