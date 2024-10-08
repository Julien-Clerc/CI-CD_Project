package tests

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"back/models"
	"back/utils"
)

func TestUserEndpoints(t *testing.T) {
	r, _ := utils.SetupTestEnvironment()

	// Test POST /users
	newUser := models.User{
		Name: "Alice",
	}
	jsonData, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, "Alice", user.Name)

	// Test GET /users
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Alice", users[0].Name)

	// Test GET /users/:id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/users/"+fmt.Sprint(user.ID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var retrievedUser models.User
	json.Unmarshal(w.Body.Bytes(), &retrievedUser)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, "Alice", retrievedUser.Name)

	// Test GET /users/:id with non-existing ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/users/9999", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var errorResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.Equal(t, "User not found", errorResponse["error"])
}
