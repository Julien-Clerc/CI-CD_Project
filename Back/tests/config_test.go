package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"back/models"
	"back/routes"
)

func TestConfigEndpoints(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	db.AutoMigrate(&models.Config{})

	r := gin.Default()

	routes.SetupRoutes(r, db)

	// Test GET /config
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/config", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var config models.Config
	json.Unmarshal(w.Body.Bytes(), &config)
	assert.Equal(t, 0, config.NbStudents)
	assert.Equal(t, 0, config.NbGroups)
	assert.Equal(t, 0, config.LastMin)
	assert.Equal(t, 0, config.LastMax)

 	// Test PUT /config
    updatedConfig := models.Config{
        NbStudents: 150,
        NbGroups:    15,
        LastMin:    10,
        LastMax:    20,
    }
    jsonData, _ := json.Marshal(updatedConfig)
    req, _ = http.NewRequest("PUT", "/config", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    json.Unmarshal(w.Body.Bytes(), &config)
    assert.Equal(t, 150, config.NbStudents)
    assert.Equal(t, 15, config.NbGroups)
    assert.Equal(t, 10, config.LastMin)
    assert.Equal(t, 20, config.LastMax)
}
