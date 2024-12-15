package userHandler

import (
	"bytes"
	"encoding/json"
	"main/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, Handler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	validator := NewValidator()
	handler := NewHandler(validator)
	router.POST("/check", handler.AddUser)
	return router, handler
}

func TestAddUser(t *testing.T) {
	router, _ := setupTestRouter()

	tests := []struct {
		name           string
		payload        models.User
		expectedStatus int
	}{
		{
			name: "Valid User",
			payload: models.User{
				Name:   "Hitesh Madgulkar",
				PAN:    "ABCDE1234F",
				Mobile: "9049871478",
				Email:  "hitesh@example.com",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid PAN",
			payload: models.User{
				Name:   "hitesh madgulkar",
				PAN:    "invalid pan",
				Mobile: "9234567890",
				Email:  "hitesh@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Mobile",
			payload: models.User{
				Name:   "hitesh madgulkar",
				PAN:    "ABCDE1234F",
				Mobile: "123", // too short
				Email:  "hitesh@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid Email",
			payload: models.User{
				Name:   "hitesh madgulkar",
				PAN:    "ABCDE1234F",
				Mobile: "1234567890",
				Email:  "invalid-email",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Required Fields",
			payload: models.User{
				Name: "hitesh madgulkar",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/check", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestAddUserInvalidJSON(t *testing.T) {
	router, _ := setupTestRouter()

	// Send invalid JSON
	invalidJSON := []byte(`{"name": "hitesh madgulkar", "pan": "ABCDE1234F", "mobile": "1234567890", "email":}`)
	req, _ := http.NewRequest("POST", "/check", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
