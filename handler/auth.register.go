package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	
}
