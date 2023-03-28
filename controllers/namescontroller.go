package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hculpan/mythictools/services"
)

func NamesPage(c *gin.Context) {
	malenames := services.GenerateListSaxonMaleNames(10)
	femalenames := services.GenerateListSaxonFemaleNames(10)

	c.HTML(http.StatusOK, "names.gohtml", gin.H{
		"malenames":   malenames,
		"femalenames": femalenames,
	})
}
