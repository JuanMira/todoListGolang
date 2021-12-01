package user

import (
	"time"

	"github.com/JuanMira/todolist/models"
	"github.com/JuanMira/todolist/utilities"
	"github.com/gofiber/fiber/v2"
)

func RegisterController(c *fiber.Ctx) error {

	db := utilities.Connect()

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&models.Error{
				Message: "Something went wrong try again",
				Err:     err.Error(),
				Status:  false,
			},
		)
	}

	password, err := utilities.EncryptPassword([]byte(user.Password))

	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(
			&models.Error{
				Message: "Something went wrong try again",
				Err:     err.Error(),
				Status:  false,
			},
		)
	}

	result := db.Create(&models.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  password,
		CreatedAt: time.Now(),
	})

	if result.RowsAffected < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			&models.Error{
				Message: "Could not create the account",
				Err:     result.Error,
				Status:  false,
			},
		)
	}

	return c.Status(fiber.StatusAccepted).JSON(
		&models.Success{
			Message: "Account created successfully",
			Status:  true,
			Data:    "",
		},
	)
}
