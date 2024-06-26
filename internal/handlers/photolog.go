package handlers

import (
	"fmt"
	"net/http"
	"plog-backend/internal/db"
	"plog-backend/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type PhotologRequest struct {
	UserID        uint           `json:"user_id"`
	GeneratedText string         `json:"generated_text"`
	Images        pq.StringArray `json:"images"`
}

func partitionExists(tableName string) bool {
	var count int64
	db.DB.Raw("SELECT count(*) FROM information_schema.tables WHERE table_name = ?", tableName).Scan(&count)
	return count > 0
}

func createPartition(tableName string, userID uint) {
	sql := fmt.Sprintf("CREATE TABLE %s PARTITION OF photologs FOR VALUES IN (%d)", tableName, userID)
	db.DB.Exec(sql)
}

func CreatePhotolog(c *gin.Context) {
	var req PhotologRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザーが存在するか確認
	var user models.User
	if err := db.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// パーティションの存在を確認し、なければ作成
	partitionTableName := fmt.Sprintf("photologs_user_%d", req.UserID)
	if !partitionExists(partitionTableName) {
		createPartition(partitionTableName, req.UserID)
	}

	photolog := models.Photolog{
		UserID:        req.UserID,
		GeneratedText: req.GeneratedText,
		Images:        req.Images,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := db.DB.Table(partitionTableName).Create(&photolog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photolog"})
		return
	}

	c.JSON(http.StatusOK, photolog)
}

func GetPhotologs(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	var photologs []models.Photolog
	partitionTableName := fmt.Sprintf("photologs_user_%s", userID)
	if err := db.DB.Table(partitionTableName).Find(&photologs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load photologs"})
		return
	}

	c.JSON(http.StatusOK, photologs)
}

func DeletePhotolog(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	photologID := c.Query("id")
	if photologID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "photolog_id is required"})
		return
	}

	partitionTableName := fmt.Sprintf("photologs_user_%s", userID)
	id, err := strconv.Atoi(photologID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid photolog_id"})
		return
	}

	if err := db.DB.Table(partitionTableName).Delete(&models.Photolog{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photolog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photolog deleted successfully"})
}
