package controllers

import "github.com/gofiber/fiber/v2"

func Index(a *fiber.App) {
	a.Get("/", index)
}

func index(c *fiber.Ctx) error {
	return c.Render("index", nil, "fragments/layout")
}
