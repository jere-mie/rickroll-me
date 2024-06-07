package src

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func setupPages(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/rickrolled", func(c *fiber.Ctx) error {
		return c.Render("rickrolled", fiber.Map{})
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		return c.Render("admin", fiber.Map{})
	})

	app.Get("/l/:urlEnding", func(c *fiber.Ctx) error {
		urlEnding := c.Params("urlEnding")

		// Find the Rickroll by URLEnding
		var rickroll Rickroll
		if err := db.Where("url_ending = ?", urlEnding).First(&rickroll).Error; err != nil {
			// just return rickroll page without passing in any data
			return c.Render("redir", fiber.Map{})
		}

		// Increment rickroll.Clicks
		rickroll.Clicks++

		// Save the updated rickroll object back to the database
		if err := db.Save(&rickroll).Error; err != nil {
			log.Errorf("Failed to update clicks for Rickroll id=%u", rickroll.ID)
			return c.Render("redir", fiber.Map{})
		}

		return c.Render("redir", fiber.Map{
			"Rickroll": rickroll,
		})
	})
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		// Read the favicon.ico file
		favicon, err := os.ReadFile("./static/img/favicon.ico")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not read favicon.ico")
		}

		// Set the appropriate content type
		c.Set("Content-Type", "image/x-icon")

		// Send the file contents as the response
		return c.Send(favicon)
	})
}
