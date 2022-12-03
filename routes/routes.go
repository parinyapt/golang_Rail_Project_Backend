package routes

import (
	// "net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/Rail_Project_Backend/handler"
	"github.com/parinyapt/Rail_Project_Backend/middleware"
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
			handler.SetupAuthAPI(v1)

			v1.Use(middleware.JWTAuth)
			{
				handler.SetupTestDBAPI(v1)
				handler.SetupStationAPI(v1)
			}
		}
	}

	// api := router.Group("/api")
	// {
	// 	v1 := api.Group("/v1")
	// 	{
	// 		v1.Use(middleware.JWTAuth)
	// 		{
	// 			handler.SetupTestDBAPI(v1)
	// 		}
	// 	}
	// }

	// auth := router.Group("/auth")
	// {
	// 	v1 := auth.Group("/v1")
	// 	{
	// 		handler.SetupAuthAPI(v1)
	// 	}
	// }

	s.ListenAndServe()
}