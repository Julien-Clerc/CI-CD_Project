package tests

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"

    "back/models"
    "back/utils"
)

func TestGroupEndpoints(t *testing.T) {
    // Set up the test environment
    r, _:= utils.SetupTestEnvironment()

    // POST /groups
    newGroup := models.Group{}
    jsonData, _ := json.Marshal(newGroup)
    req, _ := http.NewRequest("POST", "/groups", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
    var group models.Group
    json.Unmarshal(w.Body.Bytes(), &group)

    // GET /groups
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/groups", nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var groups []models.Group
    json.Unmarshal(w.Body.Bytes(), &groups)
    assert.Equal(t, 1, len(groups))

    // GET /groups/:id
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/groups/"+fmt.Sprint(group.ID), nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var retrievedGroup models.Group
    json.Unmarshal(w.Body.Bytes(), &retrievedGroup)
    assert.Equal(t, group.ID, retrievedGroup.ID)

    // DELETE /groups/:id
    w = httptest.NewRecorder()
    req, _ = http.NewRequest("DELETE", "/groups/"+fmt.Sprint(group.ID), nil)
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "Group deleted", response["message"])
}
