package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func register(c *gin.Context) {
	var reqRegister models.RequestRegister

	if err := c.ShouldBindJSON(&reqRegister); err != nil {
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
	if err := validate.Struct(reqRegister); err != nil {
		var allValidateError []models.ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			jsonfieldname, err2 := utils.GetStructTag(models.ParameterGetStructTag{
				Structx:   reqRegister,
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

	type QueryResCheckOTP struct {
		AccountUUID string
	}
	var qrCheckOTP []QueryResCheckOTP
	checkotp, err := database.DB.Raw("SELECT crp_account.account_uuid as accountuuid FROM `crp_otp` INNER JOIN crp_account ON crp_otp.otp_account_uuid = crp_account.account_uuid WHERE crp_account.account_name = '' AND `otp_uuid` = ? AND `otp_status` = '1' AND ? <= DATE_ADD(`otp_create_at`, INTERVAL 5 MINUTE)", reqRegister.RefID, time.Now()).Rows()
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
	defer checkotp.Close()
	var qrcotp QueryResCheckOTP
	for checkotp.Next() {
		checkotp.Scan(&qrcotp.AccountUUID)
		qrCheckOTP = append(qrCheckOTP, qrcotp)
	}

	if len(qrCheckOTP) != 1 {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 401,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "register fail",
				ErrorCode: "1",
				Data:      nil,
			},
		})
		return
	}

	if updatename := database.DB.Where(&models.Account{UUID: qrCheckOTP[0].AccountUUID}).Updates(models.Account{Name: reqRegister.Name}); updatename.Error != nil {
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
			Message:   "register success",
			ErrorCode: "0",
			Data:      nil,
		},
	})
}
