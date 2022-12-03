package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func routeDetail(c *gin.Context) {
	var ResponseRouteDetailData models.ResponseRouteDetailData
	sqlcommand := `SELECT
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
	row := database.DB.Raw(sqlcommand, c.GetString("language"), c.GetString("language"), c.Param("RouteID")).Row()
	if err := row.Scan(&ResponseRouteDetailData.FromLinePlatform, &ResponseRouteDetailData.FromLineName, &ResponseRouteDetailData.FromStationCode, &ResponseRouteDetailData.FromStationName, &ResponseRouteDetailData.ToLinePlatform, &ResponseRouteDetailData.ToLineName, &ResponseRouteDetailData.ToStationCode, &ResponseRouteDetailData.ToStationName, &ResponseRouteDetailData.MinPrice, &ResponseRouteDetailData.MinTime, &ResponseRouteDetailData.MinStation); err != nil {
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

	var ResponseRouteDetail []models.ResponseRouteDetail
	sqlcommand3 := `SELECT DISTINCT
	crp_line.line_platform,
	crp_route_detail.crd_station_code,
	crp_translation.translation_text
FROM
	crp_route
INNER JOIN crp_route_detail ON crp_route.route_id = crp_route_detail.crd_route_id
INNER JOIN crp_station ON crp_route_detail.crd_station_code = crp_station.station_code
INNER JOIN crp_line ON crp_route_detail.crd_line_id = crp_line.line_id
INNER JOIN crp_translation ON crp_station.station_name_tx_id = crp_translation.translation_id
INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id
WHERE
	crp_language.language_code = ? AND crp_route.route_id = ?;`
	query, err := database.DB.Raw(sqlcommand3, c.GetString("language"), c.Param("RouteID")).Rows()
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
		var rrd models.ResponseRouteDetail
		if err := query.Scan(&rrd.LinePlatform, &rrd.StationCode, &rrd.StationName); err != nil {
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
		ResponseRouteDetail = append(ResponseRouteDetail, rrd)
	}
	ResponseRouteDetailData.StationList = ResponseRouteDetail

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      ResponseRouteDetailData,
		},
	})
}
