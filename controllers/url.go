package controllers

import (
	"github.com/NickNaskida/GoTrim/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UrlController is the controller for the url resource
type UrlController struct{}

var urlService = new(services.UrlShortener)

func (u *UrlController) GetUrls(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUrls",
	})
}

func (u *UrlController) GetUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUrl",
	})
}

func (u *UrlController) CreateUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CreateUrl",
	})
}

func (u *UrlController) DeleteUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteUrl",
	})
}
