package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"back/models"
)

func GetGroups(c *gin.Context, db *gorm.DB) {
	var groups []models.Group
	db.Find(&groups)

	c.JSON(http.StatusOK, groups)
}

func GetGroup(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var group models.Group

	if err := db.First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group);
}

func CreateGroup(c *gin.Context, db *gorm.DB) {
	var group models.Group

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&group)
	c.JSON(http.StatusCreated, group)
}
