package main

import "github.com/gofiber/fiber/v2"

func HandleIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HandleCreateShortURL(c *fiber.Ctx) error {
	return c.SendString("Create Short URL")
}

func HandleRedirect(c *fiber.Ctx) error {
	return c.SendString("Redirect")
}
