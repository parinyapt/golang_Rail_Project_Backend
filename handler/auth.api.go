package handler

import (
	// "fmt"
	// "fmt"
	// "math/rand"
	"strconv"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/pkg/errors"

	// "gorm.io/gorm"

	// "github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	// models "github.com/parinyapt/Rail_Project_Backend/models/utils"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func SetupAuthAPI(router *gin.RouterGroup) {
	router.GET("/requestOTP",requestOTP)
}

func requestOTP(c *gin.Context) {
	code := strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9))

	err := utils.SendMail(models.ParameterSendMail{
		Mailto: []string{"parinyapt99@gmail.com"},
		Subject: "Hello Parinya World",
		Body: "<p>Hello Parinya Termkasipanich this is your auth otp code </p><h1>" + code + "<h1>",
		BodyType: "html",
	})
	if err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 500,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "error",
				ErrorCode: "1",
				Data:      nil,
			},
		})
		return
	}
      
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "",
			ErrorCode: "0",
			Data:      []interface{}{code},
		},
	})
}