package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiRouter := router.Group("/api")
	v1Router := apiRouter.Group("/v1")
	{
		v1Router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	
	router.NoRoute(gin.WrapH(http.FileServer(http.Dir("../dist"))))
	router.Run(":3122")
}