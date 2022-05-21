package repository

import (
	"context"

	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/domain"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/entity"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepository{db: db}
}

func (r todoRepository) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoRepository.Create")
	defer span.Finish()

	db := r.db.WithContext(ctx)
	if err := db.Create(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r todoRepository) Update(ctx context.Context, todo *entity.Todo, id int32) (*entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoRepository.Create")
	defer span.Finish()

	db := r.db.WithContext(ctx)
	if err := db.Save(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r todoRepository) Delete(ctx context.Context, id int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoRepository.Delete")
	defer span.Finish()

	todo := entity.Todo{}

	db := r.db.WithContext(ctx)
	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (r todoRepository) GetAll(ctx context.Context) (*[]entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoRepository.Create")
	defer span.Finish()

	todo := &[]entity.Todo{}

	db := r.db.WithContext(ctx)
	if err := db.Find(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r todoRepository) GetByID(ctx context.Context, id int32) (*entity.Todo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "todoRepository.Create")
	defer span.Finish()

	todo := &entity.Todo{}

	db := r.db.WithContext(ctx)
	if err := db.Where("id = ?", id).First(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
