package user

import (
	"github.com/gofiber/fiber/v2"
	database "goo/internal/db"
)

func AddUser(c *fiber.Ctx) error {

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func AllUsers(c *fiber.Ctx) error {
	user := []User{}
	database.DB.Db.Find(&user)

	return c.Status(200).JSON(user)

}
