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
	var firstHeader string
	var secondHeader string

	nameType := c.Query("type")

	if nameType == "celtic" {
		malenames = services.GenerateListCelticMaleNames(10)
		femalenames = services.GenerateListCelticFemaleNames(10)
		typename = "Celtic Names"
		firstHeader = "Males"
		secondHeader = "Females"
	} else if nameType == "town" {
		malenames = services.GenerateTownNames(10)
		femalenames = services.GenerateTownNames(10)
		typename = "Town Names"
		firstHeader = "Towns"
		secondHeader = "Towns"
	} else {
		malenames = services.GenerateListSaxonMaleNames(10)
		femalenames = services.GenerateListSaxonFemaleNames(10)
		typename = "Saxon Names"
		firstHeader = "Males"
		secondHeader = "Females"
	}

	c.HTML(http.StatusOK, "names.gohtml", gin.H{
		"malenames":    malenames,
		"femalenames":  femalenames,
		"typename":     typename,
		"type":         nameType,
		"firstheader":  firstHeader,
		"secondheader": secondHeader,
	})
}
