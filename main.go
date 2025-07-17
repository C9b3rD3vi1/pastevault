package main

import (
	"fmt"

	"github.com/C9b3rD3vi1/pastevault/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/znbang/gofiber-layout/html"
)

func main() {
	engine := html.New("./templates", ".html")

	engine.Layout("layouts/base")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// define handler routes
	app.Get("/", handlers.SecretHandler)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting the main server", err)
	}
}
