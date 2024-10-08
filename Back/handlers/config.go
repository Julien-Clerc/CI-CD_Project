package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"back/models"
)

func GetConfig(c *gin.Context, db *gorm.DB) {
	var config models.Config
	if err := db.FirstOrCreate(&config, models.Config{NbStudents: 0, NbGroups: 0, LastMin: 0, LastMax: 0}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

func UpdateConfig(c *gin.Context, db *gorm.DB) {
	var config models.Config
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&config).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}
