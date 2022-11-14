package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupAuthAPI(router *gin.RouterGroup) {
	router.POST("/requestOTP", requestOTP)
}



// err := utils.SendMail(models.ParameterSendMail{
// 	Mailto:   []string{"parinyapt99@gmail.com"},
// 	Subject:  "Hello Parinya World",
// 	Body:     "<p>Hello Parinya Termkasipanich this is your auth otp code </p><h1>" + code + "<h1>",
// 	BodyType: "html",
// })
// if err != nil {
// 	utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
// 		ResponseCode: 500,
// 		Default: utils.ResponseDefault{
// 			Success:   false,
// 			Message:   "error",
// 			ErrorCode: "1",
// 			Data:      nil,
// 		},
// 	})
// 	return
// }

// utils.ApiDefaultResponse(c, utils.ApiDefaultResponseFunctionParameter{
// 	ResponseCode: 200,
// 	Default: utils.ResponseDefault{
// 		Success:   true,
// 		Message:   "OTP send to your email",
// 		ErrorCode: "0",
// 		Data:      []interface{}{account},
// 	},
// })
// return
