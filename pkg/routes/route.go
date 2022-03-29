package route

import (
	"github.com/gofiber/fiber/v2"
	"goo/app/controller"
)

func SetupRoute(app *fiber.App) {
	app.Get("/user", user.AllUsers)
	app.Post("/add", user.AddUser)
}
