package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/model"
)

func TodoSuccessResponse(data *model.Todo) *fiber.Map {

	todo := model.Todo{
		Name:      data.Name,
		Completed: data.Completed,
	}
	return &fiber.Map{
		"status": true,
		"data":   todo,
		"error":  nil,
	}
}

func TodosSuccessResponse(data *[]model.Todo) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func TodoErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func TodoDeleteSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "",
	}
}
