package todolists

import (
	"github.com/JuanMira/todolist/models"
	taskrepositories "github.com/JuanMira/todolist/repositories/taskRepositories"
	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {

	var todolist *models.ToDoList

	if err := c.BodyParser(&todolist); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Message: "Data sended is incorrect",
			Err:     err.Error(),
			Status:  false,
		})
	}

	err := taskrepositories.CreateToDoList(todolist)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Message: "Something went wrong while trying create a new todolist, pls try again",
			Err:     err,
			Status:  false,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(models.Success{
		Message: "todolist created successfully",
		Data:    "",
		Status:  true,
	})
}
