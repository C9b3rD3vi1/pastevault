package handlers

import (
	"log"

	"github.com/C9b3rD3vi1/pastevault/database"
	"github.com/C9b3rD3vi1/pastevault/models"
	"github.com/gofiber/fiber/v2"
)

func HomePageHandler(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{})
}

func SecretHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	secret := models.Secret{}
	if err := database.DB.First(&secret, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).Render("pages/error", fiber.Map{
			"Title": "Secret Not Found",
			"Error": "The requested secret could not be found.",
		})
	}
	return c.Render("pages/secret", fiber.Map{
		"Name":   secret.Name,
		"Secret": secret.Content,
	})
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Render("pages/404", fiber.Map{})
}

func DashboardHandler(c *fiber.Ctx) error {
	var secrets []models.Secret

	// Fetch all secrets from DB
	if err := database.DB.Find(&secrets).Error; err != nil {
		log.Println("Error fetching secrets:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch secrets")
	}

	return c.Render("pages/dashboard", fiber.Map{
		"Secrets": secrets, // use plural key to match template
	})
}

func ErrorPageHandler(c *fiber.Ctx) error {
	return c.Render("pages/error", fiber.Map{})
}

// homepage hanlder
func HandleAboutPage(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{})
}
