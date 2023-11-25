package controllers

import "github.com/gofiber/fiber/v2"

func Login(a *fiber.App) {
	a.Get("/login", login)
}

func login(c *fiber.Ctx) error {
	return c.Render("login", nil, "fragments/layout")
}
