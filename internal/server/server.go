package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jnjam6681/go-clean-architecture-rest-api/config"
	v1 "github.com/jnjam6681/go-clean-architecture-rest-api/internal/delivery/http/v1"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/repository"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/usecase"
	"github.com/jnjam6681/go-clean-architecture-rest-api/pkg/database/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.ConnectionDB(cfg)
	if err != nil {
		panic(err)
	}
	gorm.Migrate()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	defer sqlDB.Close()

	todoRepo := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := v1.NewTodoHandler(todoUseCase)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	_v1 := app.Group("/api/v1")
	todoGroup := _v1.Group("/todo")
	{
		v1.TodoRouter(todoGroup, todoHandler)
	}

	log.Fatal(app.Listen(":" + cfg.App.Port))
}
