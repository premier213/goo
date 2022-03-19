package route

import (
	"goo/internal/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {

	app.Get("/user", user.AllUsers)
	app.Post("/add", user.AddUser)

}
