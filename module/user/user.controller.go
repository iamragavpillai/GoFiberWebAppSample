package user

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	c.SendString("All Users")
	return nil
}
