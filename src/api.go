package src

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupAPI(app *fiber.App, db *gorm.DB) {
	app.Post("/admin", func(c *fiber.Ctx) error {
		var req AdminPasswordRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		adminPassword := getEnv("RRM_ADMIN_PASSWORD", "")

		// if the admin password hasn't been specified, that means we disable the admin panel
		if len(adminPassword) < 1 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Admin panel has been disabled",
			})
		}
		if req.AdminPassword != adminPassword {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Incorrect Admin Password",
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
				Clicks:    r.Clicks,
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
			Clicks:          0,
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

	app.Post("/delete", func(c *fiber.Ctx) error {
		adminPassword := getEnv("RRM_ADMIN_PASSWORD", "")
		// if the admin password hasn't been specified, that means we disable the admin panel
		if len(adminPassword) < 1 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Admin panel has been disabled",
			})
		}

		var req DeleteRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if req.AdminPassword != adminPassword {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		if err := db.Delete(&Rickroll{}, req.ID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Rickroll",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Rickroll deleted successfully",
		})
	})

}
