package config

import (
	"github.com/gin-gonic/gin"
	"github.com/hculpan/mythictools/controllers"
)

func WebRoutes(r *gin.Engine) {
	r.GET("/", controllers.RootPage)
	r.GET("/about", controllers.RootPage)
	r.GET("/names", controllers.NamesPage)
}
