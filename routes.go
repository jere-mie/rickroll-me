package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
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

		return c.Render("redir", fiber.Map{
			"Rickroll": rickroll,
		})
	})

	app.Post("/admin", func(c *fiber.Ctx) error {
		var req AdminPasswordRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		adminPassword := getEnv("RRM_ADMIN_PASSWORD", "")
		if adminPassword != adminPassword {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Incorrect Admin Passeword",
			})
		}
		// Query all Rickrolls and extract urlEndings and IDs
		var rickrolls []Rickroll
		if err := db.Find(&rickrolls).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to fetch Rickrolls",
			})
		}
		rickrollInfos := make([]RickrollInfo, len(rickrolls))
		for i, r := range rickrolls {
			rickrollInfos[i] = RickrollInfo{
				ID:        r.ID,
				URLEnding: r.URLEnding,
			}
		}

		// Return the array of RickrollInfo as JSON
		return c.JSON(fiber.Map{
			"links": rickrollInfos,
		})
	})

	app.Post("/new", func(c *fiber.Ctx) error {
		var req RickrollRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		// Check for uniqueness of URLEnding
		var existing Rickroll
		if err := db.Where("url_ending = ?", req.URLEnding).First(&existing).Error; err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "URL ending already exists",
			})
		}

		// Create a new Rickroll record
		newRickroll := Rickroll{
			URLEnding:       req.URLEnding,
			SiteTitle:       req.SiteTitle,
			SiteName:        req.SiteName,
			ImgLink:         req.ImgLink,
			SiteDescription: req.SiteDescription,
		}
		if err := db.Create(&newRickroll).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to create new Rickroll",
			})
		}

		// return the resulting
		return c.JSON(fiber.Map{
			"link": fmt.Sprintf("/l/%s", newRickroll.URLEnding),
		})
	})

}
