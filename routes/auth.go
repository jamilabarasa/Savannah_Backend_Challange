package routes

import (
	"customer-orders/handler"
	"customer-orders/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	
)

//  sets up the OAuth routes for authentication
func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/:provider", func(c *gin.Context) {
		provider := c.Param("provider")
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	router.GET("/auth/:provider/callback", func(c *gin.Context) {
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Check if the user exists and create or update them
		dbUser, err := handler.CreateOrUpdateUser(user.Name, user.Email, "customer")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or update user"})
			return
		}

		// Generate JWT
		token, err := utils.GenerateJWT(dbUser.ID, dbUser.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Logged in",
			"token":   token,
			"user":    dbUser,
		})
	})
}
