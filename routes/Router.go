package routes

import (
	"github.com/JuanMira/todolist/controllers/todolists"
	"github.com/JuanMira/todolist/controllers/user"
	"github.com/JuanMira/todolist/models"
	"github.com/JuanMira/todolist/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type router struct {
	Api fiber.App
}

func NewRouter(route fiber.App) *router {
	return &router{
		Api: route,
	}
}

func (r *router) UserRoute() {
	route := r.Api.Group("/api")
	route.Post("/login", user.LoginController)
	route.Post("/register", user.RegisterController)
}

func (r *router) TaskRoute() {
	route := r.Api.Group("/api/task")

	route.Use(basicauth.New(basicauth.Config{
		Next: func(c *fiber.Ctx) bool {
			var tokenAuthorization bool
			tokenAuthorization, utilities.UserID = utilities.CheckToken(c.Get("Authorization"))
			return tokenAuthorization
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Error{
				Message: "You don't have permissions",
				Err:     nil,
				Status:  false,
			})
		},
	}))

	route.Post("/create", todolists.CreateTask)
}
