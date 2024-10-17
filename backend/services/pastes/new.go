package pastes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	z "github.com/Oudwins/zog"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shravanasati/pasteforge/backend/crud"
	"github.com/shravanasati/pasteforge/backend/utils"
)

type PasteSettings struct {
	Language           string `json:"language" zog:"language"`
	ExpirationDuration string `json:"expiration_duration" zog:"expiration_duration"`
	ExpirationNumber   uint   `json:"expiration_number" zog:"expiration_number"`
	Visibility         string `json:"visibility" zog:"visibility"`
	Password           string `json:"password" zog:"password"`
}

var stringToDurationMap = map[string]time.Duration{
	"minutes": time.Minute,
	"hours":   time.Hour,
	"days":    time.Hour * 24,
	"months":  time.Hour * 24 * 30,
	"years":    time.Hour * 24 * 365,
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

var pasteSettingSchema = z.Struct(z.Schema{
	"language":            z.String().OneOf(pasteLanguages, z.Message("unrecognized language")).Required(),
	"expiration_duration": z.String().OneOf(expirationDurations, z.Message("unrecognized expiration duration")).Required(),
	"expiration_number":   z.Int().GT(0, z.Message("expiration number must be greater than zero")).Required(),
	"visibility":          z.String().OneOf(pasteVisibilities, z.Message("unrecognized visibility")).Required(),
	"password":            z.String(),
})

func createPasteServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "unable to create a paste at the moment, please try again later",
	})
}

func (h *Handler) NewPasteHandler(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Info("unable to read json from request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to read json from request",
		})
		return
	}

	// restore the request
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var paste NewPasteRequest
	err = c.ShouldBindJSON(&paste)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
		return
	}

	var reqBody map[string]any
	err = json.Unmarshal(bodyBytes, &reqBody)
	if err != nil {
		h.logger.Info("unable to read json from request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to read json from request",
		})
		return
	}

	settingsMap, ok := reqBody["settings"]
	fmt.Println(settingsMap)
	if ok {
		// validate settings only if they exist in the request
		// use default settings otherwise
		errMap := pasteSettingSchema.Parse(settingsMap, &(paste.Settings))
		if errMap != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errMap,
			})
			return
		}
	} else {
		paste.Settings = DefaultPasteSettings()
	}

	pasteID := utils.GenerateRandomID(8)
	var pasteExpiration pgtype.Timestamptz
	if paste.Settings.ExpirationDuration == "never" {
		pasteExpiration = pgtype.Timestamptz{InfinityModifier: pgtype.Infinity}
	} else {
		duration, ok := stringToDurationMap[paste.Settings.ExpirationDuration]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid value of expiration_duration=" + paste.Settings.ExpirationDuration})
			return
		}
		pasteExpiration = pgtype.Timestamptz{InfinityModifier: pgtype.Finite, Time: time.Now().UTC().Add(duration * time.Duration(paste.Settings.ExpirationNumber))}
	}
	pasteExpiration.Valid = true
	h.logger.Debug("new paste handler", "pasteExpiratiom", pasteExpiration)
	h.logger.Info("new paste handler", "paste", paste)

	ctx := context.Background()
	tx, err := h.db.Begin(ctx)
	if err != nil {
		createPasteServerError(c)
		return
	}
	defer tx.Rollback(ctx)
	qtx := h.pasteStore.WithTx(tx)
	err = qtx.CreatePaste(ctx, crud.CreatePasteParams{
		ID:         pasteID,
		Content:    paste.Content,
		Language:   paste.Settings.Language,
		Password:   pgtype.Text{String: utils.HashPassword(paste.Settings.Password), Valid: true},
		Visibility: paste.Settings.Visibility,
		ExpiresAt:  pasteExpiration,
	})
	if err != nil {
		h.logger.Error("unable to create paste:", "err", err.Error())
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
