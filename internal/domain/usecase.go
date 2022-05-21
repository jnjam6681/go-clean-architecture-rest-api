package domain

import (
	"context"

	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
)

type TodoUsecase interface {
	Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, todo *entity.Todo, id int32) (*entity.Todo, error)
	GetAll(ctx context.Context) (*[]entity.Todo, error)
	GetByID(ctx context.Context, id int32) (*entity.Todo, error)
}
