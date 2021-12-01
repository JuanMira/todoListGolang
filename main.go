package main

import (
	"github.com/JuanMira/todolist/routes"
	"github.com/JuanMira/todolist/utilities"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := *fiber.New()

	//Migrations
	utilities.Migrate()

	//routes
	router := routes.NewRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":        "TODOList",
			"description": "todolist app made in golang and react",
			"version":     "1.0.0",
		})
	})

	router.UserRoute()
	router.TaskRoute()

	app.Listen(":8000")
}
