package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/domain"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
)

type todoHandler struct {
	u domain.TodoUsecase
}

func NewTodoHandler(usecase domain.TodoUsecase) *todoHandler {
	return &todoHandler{
		u: usecase,
	}
}

func (t *todoHandler) CreateTodo(c fiber.Ctx) error {
	var requestBody entity.Todo

	if err := c.Bind().Body(&requestBody); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(TodoErrorResponse(err))
	}

	result, err := t.u.Create(c.UserContext(), &requestBody)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(TodoErrorResponse(err))
	}
	return c.JSON(TodoSuccessResponse(result))
}

func (t *todoHandler) DeleteTodo(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	err = t.u.Delete(c.UserContext(), int32(id))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(TodoErrorResponse(err))
	}
	return c.JSON(TodoDeleteSuccessResponse())
}

func (t *todoHandler) GetAllTodos(c fiber.Ctx) error {
	todo, err := t.u.GetAll(c.UserContext())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(TodoErrorResponse(err))
	}
	return c.JSON(TodosSuccessResponse(todo))
}

func (t *todoHandler) GetTodo(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	todo, err := t.u.GetByID(c.UserContext(), int32(id))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(TodoErrorResponse(err))
	}
	return c.JSON(TodoSuccessResponse(todo))
}

func (t *todoHandler) UpdateTodo(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var requestBody entity.Todo
	if err := c.Bind().Body(&requestBody); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(TodoErrorResponse(err))
	}

	result, err := t.u.Update(c.UserContext(), &requestBody, int32(id))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(TodoErrorResponse(err))
	}
	return c.JSON(TodoSuccessResponse(result))

}
