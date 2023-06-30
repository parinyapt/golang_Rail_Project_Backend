package main

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/Rail_Project_Backend/config"
	"github.com/parinyapt/Rail_Project_Backend/database"
	// "github.com/parinyapt/Rail_Project_Backend/environment"
	"github.com/parinyapt/Rail_Project_Backend/routes"
)

func main() {
	// environment.Setup()
	database.Connect()
	config.TimezoneSetup()

	router := gin.Default()
	routes.Setup(router)
}
