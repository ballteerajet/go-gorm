package controllers

import (
	"net/http"

	"github.com/ballteerajet/go-gorm/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm" // Import GORM
)

// Modify Login to accept a db parameter
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var loginDetails struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginDetails); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("username = ?", loginDetails.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
	}
}
