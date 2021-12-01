package routes

import (
	"github.com/JuanMira/todolist/controllers/todolists"
	"github.com/JuanMira/todolist/controllers/user"
	"github.com/gofiber/fiber/v2"
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
	route.Post("/create", todolists.CreateTask)
}
