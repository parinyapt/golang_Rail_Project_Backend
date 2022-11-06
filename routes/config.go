package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func configCors(router *gin.Engine) {
	config := cors.DefaultConfig()

	// Set Allow Origins
	config.AllowAllOrigins = true
	// config.AllowOrigins = []string{
	// 	"https://prinpt.com",
	// }

	// Set Allow Methods
	config.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}

	// Set Allow Headers
	config.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type",
	}

	// Set Allow Credentials
	config.AllowCredentials = true

	// Set Max Age
	config.MaxAge = 10 * time.Minute

	router.Use(cors.New(config))
}

func configRateLimit(router *gin.Engine) {
	limiter := rate.NewLimiter(20, 1)
	router.Use(func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	})
}

func configApi(router *gin.Engine, port string) *http.Server {
	router.MaxMultipartMemory = 8 << 20

	s := &http.Server{
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if port == "" {
		log.Fatalf("[Error]->Failed to run server because no port config")
	} else {
		s.Addr = ":" + port
		log.Println("Running on PORT : " + port)
	}

	return s
}