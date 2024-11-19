package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"customer-orders/models"
)

// Global variable for the database instance
var DB *gorm.DB


// ConnectDatabase initializes the database connection.
func ConnectDatabase() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Migrate the schema (create tables based on the User and Order structs).
    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.Order{})

    // Assign the db instance to the global variable DB.
    DB = db
}

// SetDB allows setting a custom database (e.g., for testing)
func SetDB(db *gorm.DB) {
    DB = db
}
