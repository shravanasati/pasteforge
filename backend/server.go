package main

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

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