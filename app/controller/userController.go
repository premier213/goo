package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
	"goo/app/model"
	"goo/pkg/utils"
	"goo/platform/database"
)

var dbErr *pgconn.PgError

// add new user
func AddUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	erv := utils.ValidateStruct(*user)
	if erv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(erv)
	}

	if dbc := database.DB.Db.Create(&user); dbc.Error != nil {
		errors.As(dbc.Error, &dbErr)
		return c.Status(fiber.StatusBadRequest).JSON(utils.DbError(dbErr.Code, dbErr.ConstraintName))
	}

	return c.Status(200).JSON(user)

}

// show all users
func AllUsers(c *fiber.Ctx) error {
	// user := []model.User{}
	us := []model.Mp4{}
	database.DB.Db.Table("users").Select("username,password").Find(&us)

	return c.Status(200).JSON(us)
}
