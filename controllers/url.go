package controllers

import (
	"github.com/NickNaskida/GoTrim/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UrlController is the controller for the url resource
type UrlController struct{}

type urlBody struct {
	Url string `json:"url" binding:"required"`
}

func (u *UrlController) GetUrls(c *gin.Context) {
	urls := services.GetUrlService().GetAll()

	if urls == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "no urls",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"urls":    urls,
	})
}

func (u *UrlController) GetUrl(c *gin.Context) {
	url, err := services.GetUrlService().Get(c.Param("key"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "url not found",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     url,
	})
}

func (u *UrlController) CreateUrl(c *gin.Context) {
	body := urlBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
			"detail":  err.Error(),
		})
		return
	}

	url, err := services.GetUrlService().Create(body.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid url",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     url,
	})
}

func (u *UrlController) DeleteUrl(c *gin.Context) {
	err := services.GetUrlService().Delete(c.Param("key"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "url not found",
			"detail":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
