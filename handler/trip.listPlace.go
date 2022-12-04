package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func listPlace(c *gin.Context) {
	var ResponseListPlace []models.ResponseListPlace
	sqlcommand := `SELECT DISTINCT
		crp_place.place_id,
		crp_translation.translation_text,
		crp_place.place_latitude,
		crp_place.place_longitude,
		crp_place.place_distance,
		crp_place.place_thumbnail_url
	FROM
		crp_route
	INNER JOIN crp_route_detail ON crp_route.route_id = crp_route_detail.crd_route_id
	INNER JOIN crp_place ON crp_route_detail.crd_station_code = crp_place.place_station_code
	INNER JOIN crp_translation ON crp_place.place_name_tx_id = crp_translation.translation_id
	INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id
	WHERE
		crp_language.language_code = ? AND crp_place.place_category_code = 'ATTRACTION' AND crp_route.route_id = ?;`
	query, err := database.DB.Raw(sqlcommand, c.GetString("language"), c.Param("RouteID")).Rows()
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
		var rlp models.ResponseListPlace
		if err := query.Scan(&rlp.PlaceID, &rlp.PlaceName, &rlp.PlaceLatitude, &rlp.PlaceLongitude, &rlp.PlaceDistance, &rlp.PlaceImage); err != nil {
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
		ResponseListPlace = append(ResponseListPlace, rlp)
	}
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      ResponseListPlace,
		},
	})
}
