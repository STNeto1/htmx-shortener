package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Get("/", HandleIndex)
	app.Post("/shorten", HandleCreateShortURL)
	app.Get("/:url", HandleRedirect)

	app.Listen(getPort("3000"))
}

func getPort(defaultValue string) string {
	port := os.Getenv("PORT")
	if port == "" {
		return fmt.Sprintf(":%s", defaultValue)
	}
	return fmt.Sprintf(":%s", port)
}
