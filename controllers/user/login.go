package user

import (
	"github.com/JuanMira/todolist/models"
	authRepositories "github.com/JuanMira/todolist/repositories/authRepositories"
	"github.com/JuanMira/todolist/utilities"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {

	var user *models.Credentials

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "Data sended invalid",
			Err:     err.Error(),
			Status:  false,
		})
	}

	//ok if user is found, type boolean
	//uID idUser for jwt
	ok, uID, err := authRepositories.GetUser(user.Username, user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Error during the execution",
			Err:     err.Error(),
			Status:  false,
		})
	}

	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(models.Error{
			Message: "User or password not correct",
			Err:     nil,
			Status:  false,
		})
	}

	token, refreshToken, err := utilities.JWT(uID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Cannot generate a token",
			Err:     err.Error(),
			Status:  false,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(models.Success{
		Message: "Login successfully",
		Data: map[string]string{
			"refreshToken": refreshToken,
			"accessToken":  token,
		},
		Status: true,
	})
}
