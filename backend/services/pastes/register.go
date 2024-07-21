package pastes

import (
	"github.com/gin-gonic/gin"
	"github.com/shravanasati/pasteforge/backend/crud"
)

type Handler struct {
	pasteStore *crud.Queries
}

func NewHandler(db crud.DBTX) *Handler {
	return &Handler{pasteStore: crud.New(db)}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	pasteRouter := r.Group("/paste")
	pasteRouter.GET("/:id", GetPasteHandler)
	pasteRouter.POST("/new", NewPasteHandler)
}
