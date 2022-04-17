package user

import (
	"encoding/binary"
	"errors"
	"goo/app/model"
	"goo/pkg/mail"
	"goo/pkg/utils"
	"goo/pkg/validator"
	"goo/platform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

var dbErr *pgconn.PgError

func SendVerificationCode(c *fiber.Ctx) error {
	user := new(model.Email)

	// handle body parser request
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// handle validation errors
	erv := utils.ValidateStruct(*user)
	if erv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.EmailError("email_not_valid"))
	}

	// generate code for send to email and save to database
	code := uuid.New()
	byteCode := strconv.FormatUint(binary.BigEndian.Uint64(code[:8]), 10)

	// send email verification code
	if suc := mail.SendMail(user.Email, "Verification Code", byteCode); suc != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.EmailError("not_send"))
	}

	// response to client
	return c.Status(fiber.StatusOK).JSON("Verification code sent to your email")
}

// add new user
func AddUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	erv := validator.ValidateRequest(*user)
	if erv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(erv)
	}

	if dbc := database.DB.Db.Create(&user); dbc.Error != nil {
		errors.As(dbc.Error, &dbErr)
		return c.Status(fiber.StatusBadRequest).JSON(validator.ErrorDb(dbErr.Code, dbErr.ConstraintName))
	}

	return c.Status(fiber.StatusCreated).JSON(validator.SuccessRequest("user created"))

}

// show all users
func AllUsers(c *fiber.Ctx) error {
	users := []model.Mp4{}
	if data := database.DB.Db.Table("users1").Select("username").Find(&users); data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(data.Error)
	}

	return c.Status(200).JSON(users)
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
