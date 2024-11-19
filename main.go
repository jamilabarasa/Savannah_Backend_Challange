package main

import (
	"customer-orders/config"
	"customer-orders/database"
	"customer-orders/routes"
	"customer-orders/utils"
	"github.com/gin-gonic/gin"
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

	port := config.GetEnv("PORT")
	if port == ""{
		port = "3000"
	}
	
	// r.Run(":3000")
	r.Run(":"+port)

	
}
