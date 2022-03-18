package route

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	// group of routes
	api := app.Group("/")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	// v1 group
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("v1")

	})

}
