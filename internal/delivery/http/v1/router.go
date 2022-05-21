package v1

import (
	"github.com/gofiber/fiber/v2"
)

func TodoRouter(todoGroup fiber.Router, handler *todoHandler) {
	todoGroup.Post("/create", handler.CreateTodo)
	todoGroup.Get("/list", handler.GetAllTodos)
	todoGroup.Get("/get/:id", handler.GetTodo)
	todoGroup.Delete("/delete/:id", handler.DeleteTodo)
	todoGroup.Put("/update/:id", handler.UpdateTodo)
}
