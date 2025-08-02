package database

import (
	"log"
	"os"
	"shopping-cart/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL") // ✅ Read from correct env variable
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ✅ Auto migrate all required tables
	err = DB.AutoMigrate(
		&models.Cart{}, // 👈 must come before User
		&models.User{}, // has FK to Cart
		&models.Item{},
		&models.CartItem{},
		&models.Order{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
