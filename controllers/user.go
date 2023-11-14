package controllers

import (
	"errors"
	"log"
	"net/http"
	"wiai/helpers"
	"wiai/models"

	"github.com/gofiber/fiber/v2"
)

func registerUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.CreateUser(helpers.Db, username, password)

	if err != nil {
		log.Printf("failed to create user: %+v", err)

		if errors.Is(err, models.ErrUserAlreadyExists) {
			return c.SendString("User already exists!")
		}

		return c.SendString("Error occurred while registering")
	}

	log.Println("Registered: " + user.Username)

	return c.Status(http.StatusCreated).SendString("Registered successfully!")
}

func loginUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.LoadUser(helpers.Db, username, password)

	if err != nil {
		log.Printf("failed to load user: %+v", err)

		if errors.Is(err, models.ErrInvalidCredentials) {
			return c.SendString("Invalid username or password!")
		}

		return c.SendString("Error occurred while logging in")
	}

	log.Println("Logged in: " + user.Username)

	return c.SendString("ok login")
}

func User(a *fiber.App) {
	a.Post("/user/register", registerUser)
	a.Post("/user/login", loginUser)
}