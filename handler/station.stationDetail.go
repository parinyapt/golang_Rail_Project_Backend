package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func stationDetail(c *gin.Context) {
	var resStationDetail models.ResponseStationDetail
	row := database.DB.Raw("SELECT `station_id`, `station_code`, crp_translation.translation_text, `station_location_lat`, `station_location_lng`, `station_google_map`, `station_image` FROM `crp_station` INNER JOIN crp_translation ON crp_station.station_name_tx_id = crp_translation.translation_id INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND `station_id` = ?;", c.GetString("language"), c.Param("StationID")).Row()
	if err := row.Scan(&resStationDetail.StationID, &resStationDetail.StationCode, &resStationDetail.StationName, &resStationDetail.StationLatitude, &resStationDetail.StationLongitude, &resStationDetail.StationGoogleMap, &resStationDetail.StationImage); err != nil {
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
	
	var StationExitDetail []models.StationExitDetail
	sqlcommand := "SELECT `exit_number`, crp_translation.translation_text FROM `crp_station_exit` INNER JOIN crp_station ON crp_station_exit.exit_station_id = crp_station.station_id INNER JOIN crp_translation ON crp_station_exit.exit_name_tx_id = crp_translation.translation_id INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND `station_id` = ?;"
	query, err := database.DB.Raw(sqlcommand, c.GetString("language"), c.Param("StationID")).Rows()
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
		var SED2 models.StationExitDetail
		if err := query.Scan(&SED2.ExitNumber, &SED2.ExitName); err != nil {
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
		StationExitDetail = append(StationExitDetail, SED2)
	}
	resStationDetail.StationExit = StationExitDetail

	var StationFacilityDetail []models.StationFacilityDetail
	sqlcommand2 := "SELECT `facility_icon`, crp_translation.translation_text FROM `crp_station_facility` INNER JOIN crp_facility ON crp_station_facility.csf_facility_id = crp_facility.facility_id INNER JOIN crp_station ON crp_station_facility.csf_station_id = crp_station.station_id INNER JOIN crp_translation ON crp_facility.facility_name_tx_id = crp_translation.translation_id INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND `station_id` = ?;"
	query, err = database.DB.Raw(sqlcommand2, c.GetString("language"), c.Param("StationID")).Rows()
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
		var SFD models.StationFacilityDetail
		if err := query.Scan(&SFD.FacilityIcon, &SFD.FacilityName); err != nil {
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
		StationFacilityDetail = append(StationFacilityDetail, SFD)
	}
	resStationDetail.StationFacility = StationFacilityDetail

	var StationLinkDetail []models.StationLinkDetail
	sqlcommand3 := "SELECT crp_station.station_code, ctx1.translation_text, crp_line.line_platform, ctx2.translation_text FROM `crp_station_link` INNER JOIN crp_station ON crp_station_link.link_station_id_to = crp_station.station_id INNER JOIN crp_line ON crp_station.station_line_id = crp_line.line_id INNER JOIN crp_translation ctx1 ON crp_station.station_name_tx_id = ctx1.translation_id INNER JOIN crp_translation ctx2 ON crp_line.line_name_tx_id = ctx2.translation_id INNER JOIN crp_language ON ctx1.translation_language_id = crp_language.language_id AND ctx2.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND crp_station_link.link_station_id_from = ?;"
	query, err = database.DB.Raw(sqlcommand3, c.GetString("language"), c.Param("StationID")).Rows()
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
		var SLD models.StationLinkDetail
		if err := query.Scan(&SLD.StationCode, &SLD.StationName, &SLD.LinePlatform, &SLD.LineName); err != nil {
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
		StationLinkDetail = append(StationLinkDetail, SLD)
	}
	resStationDetail.StationLink = StationLinkDetail

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      resStationDetail,
		},
	})
}
