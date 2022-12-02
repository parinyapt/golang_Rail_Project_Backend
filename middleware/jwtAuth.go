package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"

	"github.com/parinyapt/Rail_Project_Backend/models"
	"github.com/parinyapt/Rail_Project_Backend/utils"
)

type AuthorizationToken struct {
	BearerToken string `validate:"required,jwt"`
}

func JWTAuth(c *gin.Context) {
	bearertoken := c.Request.Header["Authorization"]
	if bearertoken == nil || len(strings.Split(bearertoken[0], " ")) != 2 {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: http.StatusBadRequest,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Authorization Token Required",
				ErrorCode: "AUTH01",
				Data:      nil,
			},
		})
		c.Abort()
		return
	}

	tokenString := strings.Split(bearertoken[0], " ")[1]

	var authToken AuthorizationToken
	authToken.BearerToken = tokenString
	validate := validator.New()
	if err := validate.Struct(authToken); err != nil {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: http.StatusBadRequest,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Authorization Token Invalid",
				ErrorCode: "AUTH01",
				Data:      nil,
			},
		})
		c.Abort()
		return
	}

	token, _ := jwt.ParseWithClaims(tokenString, &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGN_KEY")), nil
	})

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: http.StatusBadGateway,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Authorization Invalid Method",
				ErrorCode: "AUTH02",
				Data:      nil,
			},
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(*models.JWTCustomClaims); ok && token.Valid {
		langreq := c.Request.Header["Accept-Language"]
		c.Set("language", langreq[0])
		c.Set("AccountID", claims.VerifyID)
		c.Next()
	} else {
		utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
			ResponseCode: http.StatusUnauthorized,
			Default: utils.ResponseDefault{
				Success:   false,
				Message:   "Unauthorized",
				ErrorCode: "AUTH03",
				Data:      nil,
			},
		})
		c.Abort()
		return
	}
}
