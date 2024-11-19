package handler

import (
	"customer-orders/database"
	"customer-orders/models"
	"customer-orders/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)




func SaveOrder(c *gin.Context) {
	var order models.Order

	// Bind JSON request body to order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		// Return error if binding fails
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert the customerid to a string
	id := strconv.FormatUint(uint64(order.CustomerID), 10)

	// Check if the customer exists
	customer, err := GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Save the order to the database
	if err := database.DB.Create(&order).Error; err != nil {
		// Return error if the order was not saved
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Africas's Talking Send SMS to the customer after successful order creation
	utils.SendSMS(customer.Phone, customer.Name, "Your order has been received.")

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Order was saved successfully",
	})
}
