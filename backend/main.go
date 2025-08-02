package main

import (
	"log"
	"os"
	"shopping-cart/backend/database"
	"shopping-cart/backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	// ✅ Load variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.InitDB()
	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	log.Println("Running on port:", port)
	r.Run(":" + port)
}
