package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		/*** START URL CRUD ***/
		v1.GET("/urls", _)         // TODO
		v1.GET("/urls/:key", _)    // TODO
		v1.POST("/urls", _)        // TODO
		v1.DELETE("/urls/:key", _) // TODO

		/*** START REDIRECT ***/
		v1.GET("/:key", _) // TODO
	}
}
