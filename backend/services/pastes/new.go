package pastes

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shravanasati/pasteforge/backend/crud"
	"github.com/shravanasati/pasteforge/backend/utils"
)

type PasteSettings struct {
	Language           string `json:"language"`
	ExpirationDuration string `json:"expiration_duration"`
	ExpirationNumber   uint   `json:"expiration_number"`
	Visibility         string `json:"visibility"`
	Password           string `json:"password"`
}

var stringToDurationMap = map[string]time.Duration {
	"minutes": time.Minute,
	"hours": time.Hour,
	"days": time.Hour * 24,
	"months": time.Hour * 24 * 30,
	"year": time.Hour * 24 * 365,
}

func DefaultPasteSettings() PasteSettings {
	return PasteSettings{
		Language:           "plain",
		ExpirationDuration: "never",
		ExpirationNumber:   0,
		Visibility:         "public",
		Password:           "",
	}
}

type NewPasteRequest struct {
	Content  string        `json:"content" binding:"required"`
	Settings PasteSettings `json:"settings"`
}

// fixSettings take a PasteRequestSetting and replaces zero values with default values
func fixPasteSettings(s *PasteSettings) {
	var emptyString string
	defaultSettings := DefaultPasteSettings()

	if s.Language == emptyString {
		s.Language = defaultSettings.Language
	}
	if s.ExpirationDuration == emptyString {
		s.ExpirationDuration = defaultSettings.ExpirationDuration
	}
	if s.Visibility == emptyString {
		s.Visibility = defaultSettings.Visibility
	}
}

func createPasteServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "unable to create a paste at the moment, please try again later",
	})
}

func (h *Handler) NewPasteHandler(c *gin.Context) {
	var paste NewPasteRequest
	err := c.BindJSON(&paste)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "the json payload doesn't match",
		})
		return
	}

	fixPasteSettings(&paste.Settings)
	pasteID := utils.GenerateRandomID(8)
	var pasteExpiration pgtype.Timestamp
	if paste.Settings.ExpirationDuration == "never" {
		pasteExpiration = pgtype.Timestamp{InfinityModifier: pgtype.Infinity}
	} else {
		duration, ok := stringToDurationMap[paste.Settings.ExpirationDuration]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value of expiration_duration="+paste.Settings.ExpirationDuration})
			return
		}
		pasteExpiration = pgtype.Timestamp{Time: time.Now().Add(duration * time.Duration(paste.Settings.ExpirationNumber))}
	}

	ctx := context.Background()
	tx, err := h.db.Begin(ctx)
	if err != nil {
		createPasteServerError(c)
		return
	}
	defer tx.Rollback(ctx)
	qtx := h.pasteStore.WithTx(tx)
	err = qtx.CreatePaste(ctx, crud.CreatePasteParams{
		ID: pasteID,
		Content: paste.Content,
		Language: paste.Settings.Language,
		Password: pgtype.Text{String: paste.Settings.Password},
		Visibility: paste.Settings.Visibility,
		ExpiresAt: pasteExpiration,
	})
	if err != nil {
		slog.Error("unable to create paste:", err.Error())
		createPasteServerError(c)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		createPasteServerError(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":       pasteID,
		"message":  "ok",
		"settings": paste.Settings,
	})
}
