package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, Worldd!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
