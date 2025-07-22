package handlers

import (
	"log"
	"sort"

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
		c.Locals("toast_error", "Failed to delete secret.")
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

	message := ""
		if c.Query("deleted") == "true" {
			message = "Secret deleted successfully ✅"
		}

	// Fetch all secrets from DB
	if err := database.DB.Find(&secrets).Error; err != nil {
		log.Println("Error fetching secrets:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch secrets")
	}

	// Sort by CreatedAt DESC
    sort.Slice(secrets, func(i, j int) bool {
        return secrets[i].CreatedAt.After(secrets[j].CreatedAt)
    })

	return c.Render("pages/dashboard", fiber.Map{
		"Message": message,
		"Secrets": secrets, // use plural key to match template
		"BaseURL": "https://127.0.0.1", // or read from config/env
	})
}


func ErrorPageHandler(c *fiber.Ctx) error {
	return c.Render("pages/error", fiber.Map{})
}

// homepage hanlder
func HandleAboutPage(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{})
}
