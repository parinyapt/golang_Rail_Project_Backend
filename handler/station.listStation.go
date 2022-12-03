package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func listStation(c *gin.Context) {
	var resListStation []models.ResponseListStation
	query, err := database.DB.Raw("SELECT `station_id`, `station_code`, crp_translation.translation_text, `station_location_lat`, `station_location_lng`, `station_google_map`, `station_image` FROM `crp_station` INNER JOIN crp_translation ON crp_station.station_name_tx_id = crp_translation.translation_id INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ? AND `station_line_id` = ?;", c.GetString("language"), c.Param("LineId")).Rows()
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
		var resListStation2 models.ResponseListStation
		if err := query.Scan(&resListStation2.StationID, &resListStation2.StationCode, &resListStation2.StationName, &resListStation2.StationLatitude, &resListStation2.StationLongitude, &resListStation2.StationGoogleMap, &resListStation2.StationImage); err != nil {
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
		resListStation = append(resListStation, resListStation2)
		
		// query.Scan(&qrcotp.UUID)
		// qrCheckOTP = append(qrCheckOTP, qrcotp)
		// 	// do something
	}
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "login success",
			ErrorCode: "0",
			Data:      resListStation,
		},
	})
}
