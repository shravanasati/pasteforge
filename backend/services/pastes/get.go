package pastes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shravanasati/pasteforge/backend/utils"
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

	if paste.Password.String != "" {
		password := c.Query("password")
		if !utils.ComparePasswords(paste.Password.String, password) {
			c.JSON(http.StatusForbidden, gin.H{"error": "missing/incorrect password"})
			return
		}
	}
	c.JSON(http.StatusOK, paste)
}
