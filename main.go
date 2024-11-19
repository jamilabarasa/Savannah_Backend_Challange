package main

import (
	"github.com/gin-gonic/gin"
	"customer-orders/routes"
	"customer-orders/utils"
	"customer-orders/database"	
)

func main() {
	// Load environment variables
	// config.LoadConfig()

	// Initialize the database connection
	database.ConnectDatabase()

	// Setup OAuth provider
	utils.SetupGoth()

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	routes.AuthRoutes(r)

	routes.UserRoutes(r)

	routes.OrderRoutes(r)
	
	r.Run(":3000")
}
