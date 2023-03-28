package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{})
}

func AboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.gohtml", gin.H{})
}
