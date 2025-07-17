package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	id := uuid.New().String()
	expiresAt := time.Now().Add(time.Duration(ExpireMinutes) * time.Minute)

	if ExpireMinutes != "" {
		if dur, err := time.ParseDuration(ExpireMinutes + "m"); err == nil {
			expiresAt = time.Now().Add(dur)
		}
	}

	// Create a new secret
	secret := Secret{
		ID:        id,
		Content:   Content,
		ExpiresAt: expiresAt,
	}

	// Save the secret to the database
	if err := models.db.Save(secret); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save secret",
		})
	}
	return c.Render("success", fiber.Map{
		"Link": c.BaseURL() + "/vault/" + id,
	})
}

func GetUserSecret(c *fiber.Ctx) error {
	id := c.Params("id")
	row := db.QueryRow("SELECT content, expires_at, viewed FROM secrets WHERE id = ?", id)
}
