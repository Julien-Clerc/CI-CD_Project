package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"back/models"
	"back/utils"
)

func TestConfigEndpoints(t *testing.T) {
	r, _ := utils.SetupTestEnvironment()

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
