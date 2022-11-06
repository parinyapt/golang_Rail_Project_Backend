package utils

import (
	"github.com/gin-gonic/gin"
)

type ApiDefaultResponseFunctionParameter struct {
	ResponseCode int
	Default      ResponseDefault
}

type ResponseDefault struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Data      interface{} `json:"data"`
}

func ApiDefaultResponse(c *gin.Context, param ApiDefaultResponseFunctionParameter) {
	c.JSON(param.ResponseCode, ResponseDefault{
		Success:   param.Default.Success,
		Message:   param.Default.Message,
		ErrorCode: param.Default.ErrorCode,
		Data:      param.Default.Data,
	})
}
