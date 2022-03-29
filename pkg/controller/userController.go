package user

import (
	"errors"
	"goo/db"

	// "goo/pkg/errs"
	"goo/pkg/errs"
	"goo/pkg/model"
	"goo/pkg/validator"

	// "goo/pkg/validator"

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

	erv := validator.ValidateStruct(*user)
	if erv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(erv)
	}

	if dbc := db.DB.Db.Create(&user); dbc.Error != nil {
		errors.As(dbc.Error, &dbErr)
		return c.Status(fiber.StatusBadRequest).JSON(errs.DbError(dbErr.Code, dbErr.ConstraintName))
	}

	return c.Status(200).JSON(user)

}

func AllUsers(c *fiber.Ctx) error {
	user := []model.User{}
	db.DB.Db.Find(&user)

	return c.Status(200).JSON(user)

}
