package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"customer-orders/models"
)


var DB *gorm.DB

//*gorm.DB a pointer to gorm database instance

// ConnectDatabase initializes the database connection.
func ConnectDatabase() {
    
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Migrate the schema (create tables based on the Book struct).
    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.Order{})


    // Assign the db instance to the global variable DB.
    DB = db
}
