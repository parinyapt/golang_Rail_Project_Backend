package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/parinyapt/Rail_Project_Backend/database"
	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

type RequestCreateTrip struct {
	Name          string   `json:"name" validate:"required"`
	RouteID       string   `json:"route_id" validate:"required"`
	PlaceID       []string `json:"-"`
	PlaceIDString string   `json:"place_id" validate:"required"`
}

func createTrip(c *gin.Context) {
	var reqCreateTrip RequestCreateTrip

	if err := c.ShouldBindJSON(&reqCreateTrip); err != nil {
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
	if err := validate.Struct(reqCreateTrip); err != nil {

		var allValidateError []models.ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			jsonfieldname, err2 := utils.GetStructTag(models.ParameterGetStructTag{
				Structx:   reqCreateTrip,
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

	createTripinsertdb := models.Trip{
		AccountUUID: c.GetString("AccountID"),
		Name:        reqCreateTrip.Name,
		RouteID:     reqCreateTrip.RouteID,
	}

	if tripInsert := database.DB.Create(&createTripinsertdb); tripInsert.Error != nil {
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

	reqCreateTrip.PlaceID = strings.Split(reqCreateTrip.PlaceIDString, ",")
	tripID := createTripinsertdb.ID

	for _, element := range reqCreateTrip.PlaceID {
		tripdetail := models.TripDetail{
			TripID:  tripID,
			PlaceID: element,
		}
		if tripdetailInsert := database.DB.Create(&tripdetail); tripdetailInsert.Error != nil {
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
		// fmt.Print(index,element)
		// index is the index where we are
		// element is the element from someSlice for where we are
	}
	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
		ResponseCode: 200,
		Default: utils.ResponseDefault{
			Success:   true,
			Message:   "create success",
			ErrorCode: "0",
			Data:      nil,
		},
	})
}
