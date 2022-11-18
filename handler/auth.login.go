package handler

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"

	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

func login(c *gin.Context) {
	//input and validation
	var reqLogin models.RequestLogin

	if err := c.ShouldBindJSON(&reqLogin); err != nil {
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
	if err := validate.Struct(reqLogin); err != nil {

		var allValidateError []models.ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			jsonfieldname, err2 := utils.GetStructTag(models.ParameterGetStructTag{
				Structx:   reqLogin,
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
		UUID string
	}
	var qrCheckOTP []QueryResCheckOTP
	checkotp, err := database.DB.Raw("SELECT crp_account.account_uuid as accountuuid FROM `crp_otp` INNER JOIN crp_account ON crp_otp.otp_account_uuid = crp_account.account_uuid WHERE crp_account.account_email = ? AND `otp_code` = ? AND `otp_uuid` = ? AND `otp_status` = '1' AND ? <= DATE_ADD(`otp_create_at`, INTERVAL 5 MINUTE)", reqLogin.Email, reqLogin.OTPCode, reqLogin.RefID, time.Now()).Rows()
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
		checkotp.Scan(&qrcotp.UUID)
		qrCheckOTP = append(qrCheckOTP, qrcotp)
		// 	// do something
	}

	if len(qrCheckOTP) != 1 {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: 401,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "login fail",
				ErrorCode: "1",
				Data:      nil,
			},
		})
		return
	}

	if disableAllOTP := database.DB.Where(&models.OTP{AccountUUID: qrCheckOTP[0].UUID}).Updates(models.OTP{Status: "0"}); disableAllOTP.Error != nil {
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

	// claims := models.JWTCustomClaims{
	// 	jwt.RegisteredClaims{
	// 		// A usual scenario is to set the expiration time relative to the current time
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 		NotBefore: jwt.NewNumericDate(time.Now()),
	// 		Issuer:    "RailTrip",
	// 		// Subject:   "somebody",
	// 		// ID:        "1",
	// 		// Audience:  []string{"somebody_else"},
	// 	},
	// 	VerifyID: "jsjsj",
	// }

	claims := &models.JWTCustomClaims{
		VerifyID: qrCheckOTP[0].UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "RailTrip",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringtoken, err := token.SignedString([]byte(os.Getenv("JWT_SIGN_KEY")))
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

	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "login success",
			ErrorCode: "0",
			Data:      []interface{}{
				stringtoken,
			},
		},
	})
}
