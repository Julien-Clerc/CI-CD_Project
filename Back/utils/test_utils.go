package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"back/models"
	"back/routes"
)

func SetupTestEnvironment() (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{}, &models.Config{}, &models.Group{})

	r := gin.Default()

	routes.SetupRoutes(r, db)

	return r, db
}
