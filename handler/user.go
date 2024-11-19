package handler

import (
	"customer-orders/database"
	"customer-orders/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrUpdateUser checks if the user exists, and creates or updates the user in the database
func CreateOrUpdateUser(name string, email string, role string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// User not found, so create a new user
		user = models.User{Name: name, Email: email,Phone: "N/A",Code: "N/A", Role: role}
		if err := database.DB.Create(&user).Error; err != nil {
			return nil, err
		}
	}
	return &user, nil
}

// fetches a user by their ID
func GetUserProfile(c *gin.Context) {
	id := c.Param("id");
	user, err := GetUserByID(c,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

//updates user profile : we need phone updated to send message
func UpdateUserProfile(c *gin.Context) {
	id := c.Param("id");
	//check if user exists
	user, err := GetUserByID(c,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//bind jjson
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	database.DB.Model(&user).Updates(newUser)
	c.JSON(http.StatusOK, user)

}

// checks if a user exists by their ID
func GetUserByID(c *gin.Context,id string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}