package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func listAllRoute(c *gin.Context) {
	var resListAllRoute models.ResponseListAllRoute
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
		WHERE fromlang.language_code = ? AND tolang.language_code = ? AND route_station_code_from = ? AND route_station_code_to = ?
		GROUP BY route_station_code_from, route_station_code_to;`
	row := database.DB.Raw(sqlcommand, c.GetString("language"), c.GetString("language"), c.Param("FromStationCode"), c.Param("ToStationCode")).Row()
	if err := row.Scan(&resListAllRoute.FromLinePlatform, &resListAllRoute.FromLineName, &resListAllRoute.FromStationCode, &resListAllRoute.FromStationName, &resListAllRoute.ToLinePlatform, &resListAllRoute.ToLineName, &resListAllRoute.ToStationCode, &resListAllRoute.ToStationName, &resListAllRoute.MinPrice, &resListAllRoute.MinTime, &resListAllRoute.MinStation); err != nil {
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

	var AllRouteList []models.AllRouteList
	sqlcommand3 := `SELECT
		GROUP_CONCAT(DISTINCT crp_line.line_platform),
			route_price,
			route_time,
			route_station_count
		FROM
			crp_route
		INNER JOIN crp_route_detail ON crp_route.route_id = crp_route_detail.crd_route_id
		INNER JOIN crp_line ON crp_route_detail.crd_line_id = crp_line.line_id
		WHERE
				crp_route.route_station_code_from = ? AND crp_route.route_station_code_to = ?
		GROUP BY route_id;`
	query, err := database.DB.Raw(sqlcommand3, c.Param("FromStationCode"), c.Param("ToStationCode")).Rows()
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
		var arl models.DBStructAllRouteList
		if err := query.Scan(&arl.Platform, &arl.Price, &arl.Time, &arl.Station); err != nil {
			utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
				ResponseCode: 500,
				Default: utils.ResponseDefault{
					Success:   false,
					Message:   "Internal Server Error",
					ErrorCode: "ISE02",
					Data:      nil,
				},
			})
			return
		}
		// arl.Platform = strings.Split(arl.Platform[0], ",")
		arlreal := models.AllRouteList{
			Platform: strings.Split(arl.Platform, ","),
			Price: arl.Price,
			Time: arl.Time,
			Station: arl.Station,
		}
		AllRouteList = append(AllRouteList, arlreal)
	}
	resListAllRoute.AllRouteList = AllRouteList

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      resListAllRoute,
		},
	})
}

//SELECT crp_line.line_platform, ctx2.translation_text,crp_route.`route_station_code_to`,ctx1.translation_text, MIN(`route_price`) as minprice, MIN(`route_time`) as mintime, Min(`route_station_count`) as minstation FROM `crp_route` INNER JOIN crp_station ON `route_station_code_to` = crp_station.station_code INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id INNER JOIN crp_translation ctx1 ON crp_station.station_name_tx_id = ctx1.translation_id INNER JOIN crp_translation ctx2 ON crp_line.line_name_tx_id = ctx2.translation_id INNER JOIN crp_language ON ctx1.translation_language_id = crp_language.language_id AND ctx2.translation_language_id = crp_language.language_id WHERE crp_language.language_code = 'th' AND `route_station_code_from` = 'CEN' AND `route_station_code_to` = 'BL05' GROUP BY `route_station_code_to` ORDER BY crp_line.line_platform ASC, minstation ASC;
