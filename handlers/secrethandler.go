package handlers

import (
	"log"
	"strconv"
	"time"

	"github.com/C9b3rD3vi1/pastevault/database"
	"github.com/C9b3rD3vi1/pastevault/models"
	"github.com/C9b3rD3vi1/pastevault/utils"
	"github.com/gofiber/fiber/v2"
)

// nitialize database connection

func CreateUserSecret(c *fiber.Ctx) error {
	name := c.FormValue("name")
	content := c.FormValue("content")
	expireStr := c.FormValue("expires") // assume it's number of minutes

	if name == "" || content == "" {
		return c.Render("pages/index", fiber.Map{
			"error": "Content and Name are required",
		})
	}

	// Parse expiration time
	expireMinutes, err := strconv.Atoi(expireStr)
	if err != nil || expireMinutes <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid expiration time")
	}
	expiresAt := time.Now().Add(time.Duration(expireMinutes) * time.Minute)

	// Generate random string ID (e.g., UUID or short ID)
	id := utils.GenerateID() // you can use any ID generator like UUID or nanoid

	// Create secret
	secret := models.Secret{
		ID:        id,
		Name:      name,
		Content:   content,
		ExpiresAt: expiresAt,
	}

	// Save to DB
	if result := database.DB.Create(&secret); result.Error != nil {
		log.Printf("Failed to save secret: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save secret")
	}

	// fetch all created secrets from database
	var linkSecrets []models.Secret
	if result := database.DB.Find(&linkSecrets); result.Error != nil {
		log.Printf("Failed to fetch secrets: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch secrets")
	}


	return c.Render("pages/dashboard", fiber.Map{
		// display all secrets
		"Secrets": linkSecrets,

		"Success": "Secret created successfully",
		"Link":    c.BaseURL() + "/vault/" + id,
		"BaseURL":  c.BaseURL(),
	})
}


// Get user GetUserSecret by ID
func GetUserSecret(c *fiber.Ctx) error {
	id := c.Params("id")

	var secret models.Secret
	// Look up secret by ID
	result := database.DB.First(&secret, "id = ?", id)

	// If not found or already viewed or expired
	if result.Error != nil || secret.Viewed || time.Now().After(secret.ExpiresAt) {
		return c.Status(fiber.StatusNotFound).Render("pages/error", fiber.Map{
			"Message": "This secret is unavailable or has expired.",
		})
	}
	// Mark as viewed
	secret.Viewed = true
	if err := database.DB.Save(&secret).Error; err != nil {
		log.Println("Failed to mark secret as viewed:", err)
}
	// Render the secret page
	return c.Render("pages/secret", fiber.Map{
		"Content": secret.Content,
		 "Name":    secret.Name,
	})
}

// Delete user Secret
func DeleteSecretHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	// Attempt to delete the secret by ID
	if err := database.DB.Delete(&models.Secret{}, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("pages/dashboard", fiber.Map{
			"Message": "Failed to delete secret.",
		})
	}

	// Use redirect to re-fetch updated list of secrets
	return c.Redirect("/dashboard?deleted=true", fiber.StatusSeeOther)
}

