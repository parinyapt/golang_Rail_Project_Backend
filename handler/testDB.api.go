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
	router.GET("/demo1",Demo1)
}

type ResponseDemo1 struct {
	AccountID string `json:"account_id"`
	Language string `json:"lang"`
}

func Demo1(c *gin.Context) {
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "",
			ErrorCode: "0",
			Data: ResponseDemo1{
				AccountID: c.GetString("AccountID"),
				Language: c.GetString("language"),
			},
		},
	})
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