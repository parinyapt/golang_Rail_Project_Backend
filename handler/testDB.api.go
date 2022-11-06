package handler

import (
	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func SetupTestDBAPI(router *gin.RouterGroup) {
	router.GET("/language",listLanguage)
}

func listLanguage(c *gin.Context) {
	var language models.Language

	result := database.DB.Where(&models.Language{LanguageID: 20}).First(&language)
	if result.Error != nil {
		fmt.Println("error")
		fmt.Println(result.Error)
		fmt.Println("error")
	} 

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 404,
			Default: utils.ResponseDefault{
				Success:   true,
				Message:   strconv.Itoa(int(result.RowsAffected)),
				ErrorCode: "0",
				Data:      nil,
			},
		})
		return
	}
	// print(language)
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   strconv.Itoa(int(result.RowsAffected)),
			ErrorCode: "0",
			Data:      language,
		},
	})
}