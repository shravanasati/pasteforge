package pastes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shravanasati/pasteforge/backend/utils"
)

type PasteSettings struct {
	Language           string `json:"language"`
	ExpirationDuration string `json:"expiration_duration"`
	ExpirationNumber   uint   `json:"expiration_number"`
	Visibility         string `json:"visibility"`
	Password           string `json:"password"`
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

func NewPasteHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"id":       pasteID,
		"message":  "ok",
		"settings": paste.Settings,
	})
}
