package pastes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shravanasati/pasteforge/backend/crud"
)

type Handler struct {
	db         *pgxpool.Pool
	pasteStore *crud.Queries
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{db: db, pasteStore: crud.New(db)}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	pasteRouter := r.Group("/paste")
	pasteRouter.GET("/:id", h.GetPasteHandler)
	pasteRouter.POST("/new", h.NewPasteHandler)
}
