package todolists

import (
	"github.com/JuanMira/todolist/models"
	taskrepositories "github.com/JuanMira/todolist/repositories/taskRepositories"
	"github.com/gofiber/fiber/v2"
)

//controller function to get all levels priority
func GetLevelsPriority(c *fiber.Ctx) error {

	data, err := taskrepositories.GetLevelsPriority()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Cannot get data requested",
			Err:     err.Error(),
			Status:  false,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(models.Success{
		Message: "Data fetch successfully",
		Status:  true,
		Data:    data,
	})
}
