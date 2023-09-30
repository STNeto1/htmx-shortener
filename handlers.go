package main

import "github.com/gofiber/fiber/v2"

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{}, "layouts/main")
}

func HandleCreateShortURL(c *fiber.Ctx) error {
	return c.SendString("Create Short URL")
}

func HandleRedirect(c *fiber.Ctx) error {
	return c.SendString("Redirect")
}
