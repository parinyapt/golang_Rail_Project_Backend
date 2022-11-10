package routes

import (
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/Rail_Project_Backend/handler"


)

func Setup(router *gin.Engine) {
	configCors(router)
	configRateLimit(router)
	s := configApi(router, os.Getenv("PORT"))

	//setup all api route
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			handler.SetupTestDBAPI(v1)
			handler.SetupAuthAPI(v1)
		}
	}

	s.ListenAndServe()
}