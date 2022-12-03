package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func listLine(c *gin.Context) {
	var resListLine []models.ResponseListLine
	query, err := database.DB.Raw("SELECT `line_id`, `line_platform`, `line_color, crp_translation.translation_text FROM `crp_line` INNER JOIN crp_translation ON crp_line.line_name_tx_id = crp_translation.translation_id INNER JOIN crp_language ON crp_translation.translation_language_id = crp_language.language_id WHERE crp_language.language_code = ?;", c.GetString("language")).Rows()
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
		var resLL models.ResponseListLine
		if err := query.Scan(&resLL.LineID, &resLL.LinePlatform, &resLL.LineColor, &resLL.LineName); err != nil {
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
		resListLine = append(resListLine, resLL)
	}
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "success",
			ErrorCode: "0",
			Data:      resListLine,
		},
	})
}
