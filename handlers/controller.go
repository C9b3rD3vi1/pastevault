package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HomePageHandler(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{})
}

func SecretHandler(c *fiber.Ctx) error {
	return c.Render("pages/secret", fiber.Map{})
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Render("pages/404", fiber.Map{})
}

func DashboardHandler(c *fiber.Ctx) error {
	return c.Render("pages/dashboard", fiber.Map{})
}

func ErrorPageHandler(c *fiber.Ctx) error {
	return c.Render("pages/error", fiber.Map{})
}

func HandleAboutPage(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{})
}
