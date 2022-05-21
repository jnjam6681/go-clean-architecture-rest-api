package main

import (
	"log"

	"github.com/jnjam6681/go-clean-architecture-rest-api/config"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	server.Run(cfg)
}
