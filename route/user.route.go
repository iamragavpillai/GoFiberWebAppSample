package route

import (
	"github.com/gofiber/fiber/v2"
	"main/module/user"
)

func UserRoute(app *fiber.App) {
	app.Get("/user", user.GetUsers)
}
