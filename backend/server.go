package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// env vars
var (
	PORT string
	GIN_MODE string
	DIST_DIR string
)

func validateNotEmpty(key, value string) {
	if value == "" {
		panic(key + " env var not configured")
	}
}

func validatePort(val string) {
	matches, err := regexp.MatchString(`^\d{4,5}$`, val)
	if err != nil {
		panic("error validating port" +  err.Error())
	}
	if !matches {
		panic(fmt.Sprintf("env var PORT=%s is incorrect", val))
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load env variables")
	}

	PORT = os.Getenv("PORT")
	validateNotEmpty("PORT", PORT)
	validatePort(PORT)

	GIN_MODE = os.Getenv("GIN_MODE")
	validateNotEmpty("GIN_MODE", GIN_MODE)

	DIST_DIR = os.Getenv("DIST_DIR")
	validateNotEmpty("DIST_DIR", DIST_DIR)
}

func main() {
	gin.SetMode(GIN_MODE)
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
	router.Run(":" + PORT)
}