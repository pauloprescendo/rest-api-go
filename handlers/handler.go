package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Welcome to the Home page!", "data": nil})
}
