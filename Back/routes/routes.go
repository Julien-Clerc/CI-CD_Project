package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"back/handlers"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/config", func(c *gin.Context) { handlers.GetConfig(c, db)})
	r.PUT("/config", func(c *gin.Context) { handlers.UpdateConfig(c, db)})

	r.GET("/users", func(c *gin.Context) { handlers.GetUsers(c, db)})
	r.GET("/users/:id", func(c *gin.Context) { handlers.GetUser(c, db)})
	r.POST("/users", func(c *gin.Context) { handlers.CreateUser(c, db)})
}
