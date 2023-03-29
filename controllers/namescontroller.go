package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hculpan/mythictools/services"
)

func NamesPage(c *gin.Context) {
	var malenames []string
	var femalenames []string
	var typename string

	nameType := c.Query("type")

	if nameType == "celtic" {
		malenames = services.GenerateListCelticMaleNames(10)
		femalenames = services.GenerateListCelticFemaleNames(10)
		typename = "Celtic Names"
	} else {
		malenames = services.GenerateListSaxonMaleNames(10)
		femalenames = services.GenerateListSaxonFemaleNames(10)
		typename = "Saxon Names"
	}

	c.HTML(http.StatusOK, "names.gohtml", gin.H{
		"malenames":   malenames,
		"femalenames": femalenames,
		"typename":    typename,
		"type":        nameType,
	})
}
