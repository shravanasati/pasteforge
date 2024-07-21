package misc

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/ping", PingHandler)
}
