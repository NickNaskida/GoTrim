package api

import (
	"github.com/NickNaskida/GoTrim/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		/*** START URL CRUD ***/
		url := new(controllers.UrlController)

		v1.GET("/urls", url.GetUrls)
		v1.GET("/urls/:key", url.GetUrl)
		v1.POST("/urls", url.CreateUrl)
		v1.DELETE("/urls/:key", url.DeleteUrl)
	}

	r := router.Group("/")
	{
		/*** START REDIRECT ***/
		redirect := new(controllers.RedirectController)

		r.GET("/:key", redirect.RedirectUrl)
	}

	return router
}
