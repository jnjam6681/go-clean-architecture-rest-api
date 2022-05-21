package usecase

import (
	"context"

	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/domain"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
	"github.com/opentracing/opentracing-go"
)

type todoUsecase struct {
	r domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		r: repo,
	}
}

func (u *todoUsecase) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoUsecase.Create")
	defer span.Finish()

	return u.r.Create(ctx, todo)
}

func (u *todoUsecase) Delete(ctx context.Context, id int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoUsecase.Create")
	defer span.Finish()

	return u.r.Delete(ctx, id)
}

func (u *todoUsecase) Update(ctx context.Context, todo *entity.Todo, id int32) (*entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoUsecase.Create")
	defer span.Finish()

	return u.r.Update(ctx, todo, id)
}

func (u *todoUsecase) GetAll(ctx context.Context) (*[]entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoUsecase.Create")
	defer span.Finish()

	return u.r.GetAll(ctx)
}

func (u *todoUsecase) GetByID(ctx context.Context, id int32) (*entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoUsecase.Create")
	defer span.Finish()

	return u.r.GetByID(ctx, id)
}
