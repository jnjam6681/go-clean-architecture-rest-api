package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/jnjam6681/go-clean-architecture-rest-api/config"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/delivery/http"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/repository"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/usecase"
	"github.com/jnjam6681/go-clean-architecture-rest-api/pkg/database/gorm"
)

func Run(cfg *config.Config) {
	db, err := gorm.NewGormClient(cfg)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	// postgres.SetMaxIdleConns()
	// postgres.SetConnMaxLifetime()
	// postgres.SetConnMaxIdleTime()
	// postgres.SetConnMaxIdleTime()

	defer sqlDB.Close()
	gorm.RunMigrate(db)

	todoRepo := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := http.NewTodoHandler(todoUseCase)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v1 := app.Group("/api/v1")
	todoGroup := v1.Group("/todo")

	http.TodoRouter(todoGroup, todoHandler)

	go func() {
		if err := app.Listen(":" + cfg.App.Port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	<-shutdownSignal

	log.Println("Shutdown signal received, gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server gracefully stopped")
}
