package pastes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPasteHandler(c *gin.Context) {
	pasteID := c.Param("id")
	h.logger.Debug("get paste handler", "pasteID", pasteID)
	if pasteID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing paste id in url params"})
		return
	}
	paste, err := h.pasteStore.GetPaste(context.Background(), pasteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, paste)
}