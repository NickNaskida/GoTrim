package main

import (
	"github.com/NickNaskida/GoTrim/controllers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		/*** START URL CRUD ***/
		url := new(controllers.UrlController)

		v1.GET("/urls", url.GetUrls)           // TODO
		v1.GET("/urls/:key", url.GetUrl)       // TODO
		v1.POST("/urls", url.CreateUrl)        // TODO
		v1.DELETE("/urls/:key", url.DeleteUrl) // TODO

		/*** START REDIRECT ***/
		//v1.GET("/:key", _) // TODO
	}

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
