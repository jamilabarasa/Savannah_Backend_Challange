package routes

import (
	"github.com/gin-gonic/gin"

	"customer-orders/handler"

)

func OrderRoutes(r *gin.Engine) {
	// Group routes that require authentication
	order := r.Group("/orders")
	{
		

		//TODO uncomment to Protect these routes with AuthMiddleware
		// order.Use(middleware.AuthMiddleware())


		order.POST("/", handler.SaveOrder)

		
	}
}
