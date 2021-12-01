package todolists

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "works",
	})
}
