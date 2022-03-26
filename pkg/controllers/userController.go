package user

import (
	"github.com/gofiber/fiber/v2"
	"goo/db"
	"goo/pkg/models"
)

func AddUser(c *fiber.Ctx) error {

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func AllUsers(c *fiber.Ctx) error {
	user := []models.User{}
	db.DB.Db.Find(&user)

	return c.Status(200).JSON(user)

}
