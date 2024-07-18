package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		},
	)
}

type PasteRequestSettings struct {
	Language           string `json:"language"`
	ExpirationDuration string `json:"expiration_duration"`
	ExpirationNumber   uint   `json:"expiration_number"`
	Visibility         string `json:"visibility"`
	Password           string `json:"password"`
}

type NewPasteRequest struct {
	Content  string               `json:"content" binding:"required"`
	Settings PasteRequestSettings `json:"settings"`
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

	// todo set default settings

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"settings": paste.Settings,
	})
}