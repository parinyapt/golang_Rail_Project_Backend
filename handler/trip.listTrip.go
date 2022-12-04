package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func listTrip(c *gin.Context) {
	var response []models.ResponseListTrip

	sqlcommand := `SELECT
	crp_trip.trip_id,
	crp_trip.trip_name,
	fromline.line_platform AS from_line_platform,
	fromlinets.translation_text AS from_line_name,
	fromstation.station_code AS from_station_code,
	fromstationts.translation_text AS from_station_name,
	toline.line_platform AS to_line_platform,
	tolinets.translation_text AS to_line_name,
	tostation.station_code AS to_station_code,
	tostationts.translation_text AS to_station_name,
	MIN(route_price) AS minprice,
	MIN(route_time) AS mintime,
	MIN(route_station_count) AS minstation
FROM
	crp_trip
INNER JOIN crp_route ON crp_trip.trip_route_id = crp_route.route_id
INNER JOIN crp_station fromstation ON
	route_station_code_from = fromstation.station_code
INNER JOIN crp_station tostation ON
	route_station_code_to = tostation.station_code
INNER JOIN crp_line fromline ON
	fromstation.station_line_id = fromline.line_id
INNER JOIN crp_line toline ON
	tostation.station_line_id = toline.line_id
INNER JOIN crp_translation fromstationts ON
	fromstation.station_name_tx_id = fromstationts.translation_id
INNER JOIN crp_translation fromlinets ON
	fromline.line_name_tx_id = fromlinets.translation_id
INNER JOIN crp_language fromlang ON
	fromstationts.translation_language_id = fromlang.language_id AND fromlinets.translation_language_id = fromlang.language_id
INNER JOIN crp_translation tostationts ON
	tostation.station_name_tx_id = tostationts.translation_id
INNER JOIN crp_translation tolinets ON
	toline.line_name_tx_id = tolinets.translation_id
INNER JOIN crp_language tolang ON
	tostationts.translation_language_id = tolang.language_id AND tolinets.translation_language_id = tolang.language_id
WHERE
	fromlang.language_code = ? AND tolang.language_code = ? AND crp_trip.trip_account_uuid = ?
GROUP BY
	crp_trip.trip_id;`
	query, err := database.DB.Raw(sqlcommand, c.GetString("language"), c.GetString("language"), c.GetString("AccountID")).Rows()
	if err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 500,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Internal Server Error",
				ErrorCode: "ISE01",
				Data:      nil,
			},
		})
		return
	}
	defer query.Close()
	for query.Next() {
		var rlt models.ResponseListTrip
		if err := query.Scan(&rlt.TripID ,&rlt.TripName, &rlt.FromLinePlatform, &rlt.FromLineName, &rlt.FromStationCode, &rlt.FromStationName, &rlt.ToLinePlatform, &rlt.ToLineName, &rlt.ToStationCode, &rlt.ToStationName, &rlt.MinPrice, &rlt.MinTime, &rlt.MinStation); err != nil {
			utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
				ResponseCode: 500,
				Default: utils.ResponseDefault{
					Success:   false,
					Message:   "Internal Server Error",
					ErrorCode: "ISE01",
					Data:      nil,
				},
			})
			return
		}
		response = append(response, rlt)
	}

	if len(response) == 0 {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 404,
			Default: utils.ResponseDefault{
				Success:   true,
				Message:   "not found",
				ErrorCode: "0",
				Data:      nil,
			},
		})
		return
	}

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      response,
		},
	})
}
