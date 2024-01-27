package tests

import (
	"encoding/json"
	"fmt"
	"github.com/NickNaskida/GoTrim/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrlsEmpty(t *testing.T) {
	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/urls", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"success\",\"urls\":{}}", w.Body.String())
}

func TestGetUrlNotFound(t *testing.T) {
	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/urls/123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "url not found")
}

func TestCreateUrl(t *testing.T) {
	router := api.SetupRouter()

	urlToShorten := "https://google.com"
	validBody := fmt.Sprintf("{\"url\":\"%s\"}", urlToShorten)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/urls", strings.NewReader(validBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "success", response["message"])
	assert.Contains(t, response, "url")

	urlValue, ok := response["url"].(string)
	assert.True(t, ok)

	key := strings.Split(urlValue, "http://localhost:8080/")[1]
	assert.NotEmpty(t, key)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/urls/"+key, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), urlToShorten)
}

func TestCreateUrlInvalidBody(t *testing.T) {
	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/urls", strings.NewReader("{\"invalid\":\"body\"}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid body")
}

func TestCreateUrlInvalidUrl(t *testing.T) {
	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/urls", strings.NewReader("{\"url\":\"invalid\"}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid url")
}

func TestDeleteUrl(t *testing.T) {
	router := api.SetupRouter()

	urlToShorten := "https://google.com"
	validBody := fmt.Sprintf("{\"url\":\"%s\"}", urlToShorten)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/urls", strings.NewReader(validBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "success", response["message"])
	assert.Contains(t, response, "url")

	urlValue, ok := response["url"].(string)
	assert.True(t, ok)

	key := strings.Split(urlValue, "http://localhost:8080/")[1]
	assert.NotEmpty(t, key)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/urls/"+key, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), urlToShorten)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/v1/urls/"+key, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/urls/"+key, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "url not found")
}

func TestDeleteUrlNotFound(t *testing.T) {
	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/urls/123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "url not found")
}
