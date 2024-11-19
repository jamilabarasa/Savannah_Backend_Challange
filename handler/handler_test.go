package handler_test

import (
	"bytes"
	"customer-orders/models"
	"customer-orders/handler"
	"customer-orders/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func TestSaveOrder(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Create an in-memory SQLite database for testing
	sqliteDB, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	sqliteDB.AutoMigrate(&models.User{}, &models.Order{})

	// Inject the in-memory database into the handler
	database.SetDB(sqliteDB)

	// Register routes
	router.POST("/orders", handler.SaveOrder)

	t.Run("successful order save", func(t *testing.T) {
		// Create test user in the mock database
		user := models.User{
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}
		sqliteDB.Create(&user)

		// Parse the string date into a time.Time
		orderTime, err := time.Parse(time.RFC3339Nano, "2024-11-19T13:17:28.004831698+03:00")
		if err != nil {
			t.Fatalf("failed to parse time: %v", err)
		}

		// Create order request with parsed time
		order := models.Order{
			Item:       "Test Item",
			Amount:     100.50,
			Time:       orderTime,
			CustomerID: uint(user.ID),
		}
		orderJSON, _ := json.Marshal(order)

		// Send request
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(orderJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// Serve the request
		router.ServeHTTP(resp, req)

		// Assertions
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "Order was saved successfully")
	})

	t.Run("missing customer ID", func(t *testing.T) {
		// Parse the string date into a time.Time
		orderTime, err := time.Parse(time.RFC3339Nano, "2024-11-19T13:17:28.004831698+03:00")
		if err != nil {
			t.Fatalf("failed to parse time: %v", err)
		}

		// Send request with an invalid customer ID
		order := models.Order{
			Item:       "Test Item",
			Amount:     100.50,
			Time:       orderTime, 
			CustomerID: 999, 
		}
		orderJSON, _ := json.Marshal(order)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(orderJSON))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// Serve the request
		router.ServeHTTP(resp, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, resp.Code)
		assert.Contains(t, resp.Body.String(), "user not found")
	})
}
