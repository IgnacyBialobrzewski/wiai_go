package controllers

import (
	"errors"
	"log"
	"net/http"
	"wiai/helpers"
	"wiai/models"

	"github.com/gofiber/fiber/v2"
)

const sessionCookieKey = "session"

func Sessions(a *fiber.App) {
	a.Post("/sessions", createSession)
	a.Delete("/sessions", deleteSession)
}

func deleteSession(c *fiber.Ctx) error {
	sessionKey := c.Cookies(sessionCookieKey)
	err := models.DeleteSession(helpers.Db, sessionKey)

	if sessionKey == "" {
		return c.SendStatus(http.StatusUnauthorized)
	}

	c.ClearCookie(sessionCookieKey)

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Couldn't log out!")
	}

	return c.SendStatus(http.StatusOK)
}

func createSession(c *fiber.Ctx) error {
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

	session, err := models.CreateSession(helpers.Db, user)

	if err != nil {
		log.Printf("failed to create a session: %+v", err)
		return c.SendString("Session error occured")
	}

	c.Cookie(&fiber.Cookie{
		Name:   sessionCookieKey,
		Value:  session.Key,
		Secure: true,
	})

	log.Println("Logged in: " + user.Username)

	return c.SendString("Logged in successfully!")
}