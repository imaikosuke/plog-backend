package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "plog-backend/internal/db"
    "plog-backend/internal/models"
    "golang.org/x/crypto/bcrypt"
    "log"
)

func LoginUser(c *gin.Context) {
    var request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := db.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    log.Printf("Stored hash: %s", user.PasswordHash)
    log.Printf("Provided password: %s", request.Password)

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
        log.Printf("Password comparison failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
