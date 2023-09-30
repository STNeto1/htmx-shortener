package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{}, "layouts/main")
}

type CreateShortURLRequest struct {
	URL string `form:"url,required"`
}

func HandleCreateShortURL(c *fiber.Ctx) error {
	var body CreateShortURLRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Render("pages/index", fiber.Map{
			"error": err.Error(),
		})
	}

	hashedLink, err := CreateShortenedLink(body.URL)
	if err != nil {
		log.Println("failed to create shortened link", err)
		return c.Render("pages/index", fiber.Map{
			"error": "failed to create shortened link",
			"url":   body.URL,
		})
	}

	log.Println("hashed link", hashedLink)

	return c.Render("pages/index", fiber.Map{
		"message": "URL shortened!",
	}, "layouts/main")
}

func HandleRedirect(c *fiber.Ctx) error {
	return c.SendString("Redirect")
}
