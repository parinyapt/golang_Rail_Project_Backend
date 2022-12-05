package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func deleteTrip(c *gin.Context) {
	var tempDbTrip DBTrip
	sqlcommand := `SELECT
		trip_name,
		trip_route_id
	FROM
		crp_trip
	WHERE
		trip_id = ? AND trip_account_uuid = ?`
	row := database.DB.Raw(sqlcommand, c.Param("TripID"), c.GetString("AccountID")).Row()
	if err := row.Scan(&tempDbTrip.Name, &tempDbTrip.RouteID); err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 404,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Trip Not found",
				ErrorCode: "404",
				Data:      nil,
			},
		})
		return
	}

	if deleteTrip := database.DB.Where("trip_id = ? AND trip_account_uuid = ?", c.Param("TripID"), c.GetString("AccountID")).Delete(&models.Trip{}); deleteTrip.Error != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 500,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Internal Server Error",
				ErrorCode: "ISE03",
				Data:      nil,
			},
		})
		return
	}

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "Trip delete success",
			ErrorCode: "0",
			Data:      nil,
		},
	})
}
