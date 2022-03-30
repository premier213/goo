package user

import (
	"errors"
	"goo/app/model"
	"goo/pkg/utils"
	"goo/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
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

	return c.Status(fiber.StatusCreated).JSON(user)

}

// show all users
func AllUsers(c *fiber.Ctx) error {
	us := []model.Mp4{}
	database.DB.Db.Table("users").Select("username,password").Find(&us)

	return c.Status(200).JSON(us)
}

func FindUser(c *fiber.Ctx) error {
	mp4 := []model.Mp4{}
	p := new(model.Mp4)
	c.Query("username")

	if err := c.QueryParser(p); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Table("users").Where("username LIKE ?", "%"+p.Username+"%").Find(&mp4)
	return c.Status(200).JSON(mp4)
}
