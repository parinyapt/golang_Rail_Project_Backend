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

func requestOTPnew(c *gin.Context) {
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
		var allValidateError []models.ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			jsonfieldname, err2 := utils.GetStructTag(models.ParameterGetStructTag{
				Structx:   reqOTP,
				FieldName: err.Field(),
				TagName:   "json",
			})
			if err2 != nil {
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
			allValidateError = append(allValidateError, models.ValidatorError{
				Field:    jsonfieldname,
				ErrorMsg: utils.ValidatorErrorMsg(err.Tag()),
			})
		}

		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 400,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "error invalid input",
				ErrorCode: "2",
				Data:      allValidateError,
			},
		})
		return
	}

	registerStatus := "login"

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
		registerStatus = "register"
		accountRegister := models.Account{
			UUID:  uuid.New().String(),
			Email: reqOTP.Email,
			Name:  "",
		}
		if registerInsert := database.DB.Create(&accountRegister); registerInsert.Error != nil {
			utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
				ResponseCode: 500,
				Default: utils.ResponseDefault{
					Success:   false,
					Message:   "Internal Server Error",
					ErrorCode: "ISE05",
					Data:      nil,
				},
			})
			return
		}
		account = models.Account{
			UUID: accountRegister.UUID,
			Email: accountRegister.Email,
		}
	}

	var otp models.OTP
	checkOTP1 := database.DB.Where("otp_account_uuid = ? AND otp_status = ? AND ? <= DATE_ADD(otp_create_at, INTERVAL 1 MINUTE)", account.UUID, "1", time.Now()).Find(&otp)
	if checkOTP1.Error != nil {
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
				ErrorCode: "ISE03",
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
				ErrorCode: "ISE04",
				Data:      nil,
			},
		})
		return
	}

	if err := utils.SendMail(models.ParameterSendMail{
		Mailto:   []string{account.Email},
		Subject:  "RailTrip Login OTP",
		Body:     "<p>Your OTP is </p><h1>" + code + "</h1><p>(Valid for 5 minutes)</p><p>REF:" + otpinsertdata.UUID + "</p>",
		BodyType: "html",
	}); err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 500,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Internal Server Error",
				ErrorCode: "ISE06",
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
			Data: models.ResponserequestOTP{
				RefID:  otpinsertdata.UUID,
				Status: registerStatus,
			},
		},
	})
}
