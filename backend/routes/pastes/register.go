package pastes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	pasteRouter := r.Group("/paste")
	pasteRouter.POST("/new", NewPasteHandler)
}
