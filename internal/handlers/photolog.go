package handlers

import (
	"net/http"
	"plog-backend/internal/db"
	"plog-backend/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type PhotologRequest struct {
	UserID        uint     `json:"user_id"`
	GeneratedText string   `json:"generated_text"`
	Images        []string `json:"images"`
}

func CreatePhotolog(c *gin.Context) {
	var req PhotologRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photolog := models.Photolog{
		UserID:        req.UserID,
		GeneratedText: req.GeneratedText,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := db.DB.Create(&photolog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photolog"})
		return
	}

	for _, imageURL := range req.Images {
		image := models.Image{
			PhotologID: photolog.ID,
			ImageURL:   imageURL,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		if err := db.DB.Create(&image).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
	}

	c.JSON(http.StatusOK, photolog)
}
