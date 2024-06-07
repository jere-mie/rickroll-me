package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	PORT := getEnv("RRM_PORT", "3000")

	engine := django.New("./templates", ".html")

	// Create a Fiber app with the configured engine
	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
	})

	app.Use(logger.New())

	// Serve static files
	app.Static("/static", "./static")

	// Initialize SQLite database
	db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&Rickroll{}); err != nil {
		log.Fatal(err)
	}

	setupRoutes(app, db)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
