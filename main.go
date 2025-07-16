package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/znbang/gofiber-layout/html"
)

func main() {
	engine := html.New("./templates", ".html")

	engine.Layout("layout/base")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting the main server", err)
	}
}
