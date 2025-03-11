package domain

import (
	"context"

	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/model"
)

type TodoUsecase interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, todo *model.Todo, id int32) (*model.Todo, error)
	GetAll(ctx context.Context) (*[]model.Todo, error)
	GetByID(ctx context.Context, id int32) (*model.Todo, error)
}
