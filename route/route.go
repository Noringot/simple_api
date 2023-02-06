package route

import (
	"github.com/Noringotq/api/mw"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Use(mw.Auth)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, User!")
	})
}
