package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"

	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func requestOTP(c *gin.Context) {
	//input and validation
	var reqOTP models.RequestrequestOTP

	if err := c.ShouldBindJSON(&reqOTP); err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 400,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "error",
				ErrorCode: "1",
				Data:      nil,
			},
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqOTP); err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 400,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "error invalid input",
				ErrorCode: "2",
				Data:      nil,
			},
		})
		return
	}

	//find account in database
	var account models.Account
	// findAccount := database.DB.Where(&models.Account{Email: reqOTP.Email}).First(&account)
	findAccount := database.DB.Where("account_email = ?", reqOTP.Email).Find(&account)
	if findAccount.Error != nil {
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
	if findAccount.RowsAffected == 0 {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 404,
			Default: utils.ResponseDefault{
				Success:   true,
				Message:   "",
				ErrorCode: "0",
				Data:      nil,
			},
		})
		return
	}

	var otp models.OTP
	checkOTP1 := database.DB.Where("otp_account_uuid = ? AND otp_status = ? AND ? <= DATE_ADD(otp_create_at, INTERVAL 1 MINUTE)", account.UUID, "1", time.Now()).Find(&otp)
	if checkOTP1.Error != nil {
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
	if checkOTP1.RowsAffected > 0 {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 400,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "can request otp again in 1 minute",
				ErrorCode: "1",
				Data:      nil,
			},
		})
		return
	}

	if disableAllOTP := database.DB.Where(&models.OTP{AccountUUID: account.UUID}).Updates(models.OTP{Status: "0"}); disableAllOTP.Error != nil {
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

	code := strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9)) + strconv.Itoa(utils.RandomNumber(9))
	otpinsertdata := models.OTP{
		UUID:        uuid.New().String(),
		AccountUUID: account.UUID,
		Code:        code,
		Status:      "1",
	}
	if otpInsert := database.DB.Create(&otpinsertdata); otpInsert.Error != nil {
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

	if err := utils.SendMail(models.ParameterSendMail{
		Mailto:   []string{account.Email},
		Subject:  "Hello " + account.Name,
		Body:     "<p>Your OTP is </p><h1>" + code + "</h1><p>(Valid for 5 minutes). REF:" + otpinsertdata.UUID + "</p>",
		BodyType: "html",
	}); err != nil {
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

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "OTP send to your Email",
			ErrorCode: "0",
			Data:      models.ResponserequestOTP{
				RefID: otpinsertdata.UUID,
			},
		},
	})
}