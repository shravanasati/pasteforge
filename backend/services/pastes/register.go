package pastes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shravanasati/pasteforge/backend/crud"
)

type Handler struct {
	logger     *slog.Logger
	db         *pgxpool.Pool
	pasteStore *crud.Queries
}

func NewHandler(logger *slog.Logger, db *pgxpool.Pool) *Handler {
	return &Handler{logger: logger, db: db, pasteStore: crud.New(db)}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	pasteRouter := r.Group("/pastes")
	pasteRouter.GET("/get/:id", h.GetPasteHandler)
	pasteRouter.POST("/new", h.NewPasteHandler)
}
