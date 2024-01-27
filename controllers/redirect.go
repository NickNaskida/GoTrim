package controllers

import (
	"github.com/NickNaskida/GoTrim/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedirectController struct{}

func (r *RedirectController) RedirectUrl(c *gin.Context) {
	url, err := services.GetUrlService().Get(c.Param("key"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "url not found",
			"detail":  err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
