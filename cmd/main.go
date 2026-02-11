package main

import (
	"log"
	"os"
	"resign-api/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	//import class to main to called the function to connect to supabase
	db := database.NewPostgresDB()

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		sqlDB, _ := db.DB()
		err := sqlDB.Ping()

		status := "OK"
		if err != nil {
			status = "Error connecting to database"
		}

		return c.JSON(fiber.Map{
			"message":   "Fiber is running!",
			"db_status": status,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("port not set in env.")
	}

	log.Printf("Server is starting in port  %s...", port)
	log.Fatal(app.Listen(":" + port))
}
