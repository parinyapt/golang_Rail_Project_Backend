package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"

	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

type DBTrip struct {
	Name    string
	RouteID string
}

func tripDetail(c *gin.Context) {
	var response models.ResponseTripDetail

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
				Message:   "Not found",
				ErrorCode: "404",
				Data:      nil,
			},
		})
		return
	}
	response.TripName = tempDbTrip.Name

	sqlcommand = `SELECT
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
		crp_route
	INNER JOIN crp_station fromstation ON route_station_code_from = fromstation.station_code
	INNER JOIN crp_station tostation ON route_station_code_to = tostation.station_code
	INNER JOIN crp_line fromline ON fromstation.station_line_id = fromline.line_id
	INNER JOIN crp_line toline ON tostation.station_line_id = toline.line_id
	INNER JOIN crp_translation fromstationts ON fromstation.station_name_tx_id = fromstationts.translation_id
	INNER JOIN crp_translation fromlinets ON fromline.line_name_tx_id = fromlinets.translation_id
	INNER JOIN crp_language fromlang ON fromstationts.translation_language_id = fromlang.language_id AND fromlinets.translation_language_id = fromlang.language_id
	INNER JOIN crp_translation tostationts ON tostation.station_name_tx_id = tostationts.translation_id
	INNER JOIN crp_translation tolinets ON toline.line_name_tx_id = tolinets.translation_id
	INNER JOIN crp_language tolang ON tostationts.translation_language_id = tolang.language_id AND tolinets.translation_language_id = tolang.language_id
	WHERE fromlang.language_code = ? AND tolang.language_code = ? AND route_id = ?
	GROUP BY route_station_code_from, route_station_code_to;`
	row = database.DB.Raw(sqlcommand, c.GetString("language"), c.GetString("language"), tempDbTrip.RouteID).Row()
	if err := row.Scan(&response.FromLinePlatform, &response.FromLineName, &response.FromStationCode, &response.FromStationName, &response.ToLinePlatform, &response.ToLineName, &response.ToStationCode, &response.ToStationName, &response.MinPrice, &response.MinTime, &response.MinStation); err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 404,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Not found",
				ErrorCode: "404",
				Data:      nil,
			},
		})
		return
	}

	var arrayplacedetail []models.PlaceDetail
	sqlcommand = `SELECT
		crp_line.line_id,
		crp_line.line_platform,
		crp_line.line_color,
		linetx.translation_text,
		crp_station.station_id,
		crp_station.station_code,
		stationtx.translation_text,
		crp_place.place_id,
		placetx.translation_text,
		crp_place.place_latitude,
		crp_place.place_longitude,
		crp_place.place_distance,
		crp_place.place_thumbnail_url
	FROM
		crp_trip_detail
	INNER JOIN crp_place ON crp_trip_detail.ctd_place_id = crp_place.place_id
	INNER JOIN crp_station ON crp_place.place_station_code = crp_station.station_code
	INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id
	INNER JOIN crp_translation placetx ON
		crp_place.place_name_tx_id = placetx.translation_id
	INNER JOIN crp_translation stationtx ON
		crp_station.station_name_tx_id = stationtx.translation_id
	INNER JOIN crp_translation linetx ON
		crp_line.line_name_tx_id = linetx.translation_id
	INNER JOIN crp_language ON placetx.translation_language_id = crp_language.language_id AND stationtx.translation_language_id = crp_language.language_id AND linetx.translation_language_id = crp_language.language_id
	WHERE crp_language.language_code = ? AND crp_trip_detail.ctd_trip_id = ?
	GROUP BY crp_place.place_id
	ORDER BY crp_line.line_id,crp_station.station_id;`
	query, err := database.DB.Raw(sqlcommand, c.GetString("language"), c.Param("TripID")).Rows()
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
		var apd models.PlaceDetail
		if err := query.Scan(&apd.LineID, &apd.LinePlatform, &apd.LineColor, &apd.LineName, &apd.StationID, &apd.StationCode, &apd.StationName ,&apd.PlaceID, &apd.PlaceName, &apd.PlaceLatitude, &apd.PlaceLongitude, &apd.PlaceDistance, &apd.PlaceImage); err != nil {
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
		arrayplacedetail = append(arrayplacedetail, apd)
	}
	response.PlaceDetail = arrayplacedetail
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
