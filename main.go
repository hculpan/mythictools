package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hculpan/mythictools/config"
	"github.com/hculpan/mythictools/services"
)

func init() {
	config.LoadEnvVariables()
	// config.ConnectToDB()
	// config.LoadCateogires()
}

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.gohtml")

	services.InitNameGeneratorService()

	config.WebRoutes(r)

	r.Run()
}
