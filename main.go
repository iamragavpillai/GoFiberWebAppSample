package main

import (
	"github.com/gofiber/fiber/v2"
	"main/route"
)

func main() {
	app := fiber.New()

	route.UserRoute(app)

	app.Listen(":3000")
}
