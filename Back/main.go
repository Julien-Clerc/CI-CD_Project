package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"back/models"
	"back/routes"
)

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the SQLite database successfully")

	db.AutoMigrate(&models.Config{})

	r := gin.Default()

	routes.SetupRoutes(r, db)

	r.Run(":3000")
}
