package routes

import (
	"github.com/gin-gonic/gin"

	"customer-orders/handler" 
	"customer-orders/middleware"
)

func UserRoutes(r *gin.Engine) {
	// Group routes that require authentication
	user := r.Group("/users")
	{
		// Protect these routes with AuthMiddleware
		user.Use(middleware.AuthMiddleware())

		
		user.GET("/:id", handler.GetUserProfile)

		
		user.PUT("/:id", handler.UpdateUserProfile)
	}
}
