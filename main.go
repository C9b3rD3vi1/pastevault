package main

import (
	"fmt"

	"github.com/C9b3rD3vi1/pastevault/database"
	"github.com/C9b3rD3vi1/pastevault/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/znbang/gofiber-layout/html"
)

func main() {
	engine := html.New("./templates", ".html")

	engine.Layout("layouts/base")

	// Initialize database
	database.InitDB()

	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	// define handler routes
	app.Get("/", handlers.HomePageHandler)
	app.Get("/dashboard", handlers.DashboardHandler)
	app.Get("/error", handlers.ErrorPageHandler)
	app.Get("/secret/:id", handlers.CreateUserSecret)
	app.Post("/secret", handlers.CreateUserSecret)
	app.Get("/secret", handlers.SecretHandler)
	app.Get("/about", handlers.HandleAboutPage)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting the main server", err)
	}
}
