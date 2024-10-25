package controllers

import (
	"net/http"

	"github.com/ballteerajet/go-gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get User by ID
func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// Create User
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// Update User
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}

// Delete User
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.User{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
