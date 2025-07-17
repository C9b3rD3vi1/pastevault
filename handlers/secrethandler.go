package handlers

import (
	"strconv"
	"time"

	"github.com/C9b3rD3vi1/pastevault/database"
	"github.com/C9b3rD3vi1/pastevault/models"
	"github.com/gofiber/fiber/v2"
)

// nitialize database connection

func CreateUserSecret(c *fiber.Ctx) error {

	Content := c.FormValue("content")
	ExpireMinutes := c.FormValue("expires")

	if Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Content is required",
		})
	}
	// assign id and expiry time
	idStr := c.Params("id")
	expireMinutes := c.FormValue("expire_minutes") // this is a string
	minutesInt, err := strconv.Atoi(expireMinutes)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid expiration time")
	}

	expiresAt := time.Now().Add(time.Duration(minutesInt) * time.Minute)

	if ExpireMinutes != "" {
		if dur, err := time.ParseDuration(ExpireMinutes + "m"); err == nil {
			expiresAt = time.Now().Add(dur)
		}
	}

	// Convert idStr := c.Params("id") to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	// Create a new secret
	secret := models.Secret{
		ID:        id,
		Content:   Content,
		ExpiresAt: expiresAt,
	}

	// Save the secret to the database
	if err := database.DB.Save(secret); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save secret",
		})
	}
	return c.Render("success", fiber.Map{
		"Link": c.BaseURL() + "/vault/" + strconv.Itoa(id), // convert id to string agin
	})
}

func GetUserSecret(c *fiber.Ctx) error {
	id := c.Params("id")

	var secret models.Secret
	// Look up secret by ID
	result := database.DB.First(&secret, "id = ?", id)

	// If not found or already viewed or expired
	if result.Error != nil || secret.Viewed || time.Now().After(secret.ExpiresAt) {
		return c.SendString("This secret is unavailable or has expired.")
	}

	// Mark as viewed
	secret.Viewed = true
	database.DB.Save(&secret)

	// Render the secret page
	return c.Render("pages/secret", fiber.Map{
		"Content": secret.Content,
	})
}

func SecretHandler(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{})
}
