package handler

import (

	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func listToStation(c *gin.Context) {
	sort, have := c.GetQuery("sort")
	if have {
		if sort == "price" {
			sort = "minprice"
		}else if sort == "time" {
			sort = "mintime"
		}else{
			sort = "minstation"
		}
	} else {
		sort = "minstation"
	}

	var resListToStation []models.ResponseListToStation
	// fmt.Println(sort)
	query, err := database.DB.Raw("SELECT crp_line.line_platform, ctx2.translation_text,crp_route.`route_station_code_to`,ctx1.translation_text, MIN(`route_price`) as minprice, MIN(`route_time`) as mintime, Min(`route_station_count`) as minstation FROM `crp_route` INNER JOIN crp_station ON `route_station_code_to` = crp_station.station_code INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id INNER JOIN crp_translation ctx1 ON crp_station.station_name_tx_id = ctx1.translation_id INNER JOIN crp_translation ctx2 ON crp_line.line_name_tx_id = ctx2.translation_id INNER JOIN crp_language ON ctx1.translation_language_id = crp_language.language_id AND ctx2.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND `route_station_code_from` = ? GROUP BY `route_station_code_to` ORDER BY crp_line.line_platform ASC, " + sort + " ASC;", c.GetString("language"), c.Param("FromStationCode")).Rows()
	// fmt.Println(query.Columns())
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
		var rls models.ResponseListToStation
		if err := query.Scan(&rls.LinePlatform, &rls.LineName, &rls.StationCode, &rls.StationName, &rls.MinPrice, &rls.MinTime, &rls.MinStation); err != nil {
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
		resListToStation = append(resListToStation, rls)
	}
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      resListToStation,
		},
	})
}

//SELECT crp_line.line_platform, ctx2.translation_text,crp_route.`route_station_code_to`,ctx1.translation_text, MIN(`route_price`) as minimum FROM `crp_route` INNER JOIN crp_station ON `route_station_code_to` = crp_station.station_code INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id INNER JOIN crp_translation ctx1 ON crp_station.station_name_tx_id = ctx1.translation_id INNER JOIN crp_translation ctx2 ON crp_line.line_name_tx_id = ctx2.translation_id INNER JOIN crp_language ON ctx1.translation_language_id = crp_language.language_id AND ctx2.translation_language_id = crp_language.language_id WHERE crp_language.language_code = 'th' AND `route_station_code_from` = 'BL01' GROUP BY `route_station_code_to`;

//SELECT crp_line.line_platform, ctx2.translation_text,crp_route.`route_station_code_to`,ctx1.translation_text, MIN(`route_price`) as minprice, MIN(`route_time`) as mintime, Min(`route_station_count`) as minstation FROM `crp_route` INNER JOIN crp_station ON `route_station_code_to` = crp_station.station_code INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id INNER JOIN crp_translation ctx1 ON crp_station.station_name_tx_id = ctx1.translation_id INNER JOIN crp_translation ctx2 ON crp_line.line_name_tx_id = ctx2.translation_id INNER JOIN crp_language ON ctx1.translation_language_id = crp_language.language_id AND ctx2.translation_language_id = crp_language.language_id WHERE crp_language.language_code = 'th' AND `route_station_code_from` = 'BL01' GROUP BY `route_station_code_to` ORDER BY crp_line.line_platform ASC, minstation ASC;
