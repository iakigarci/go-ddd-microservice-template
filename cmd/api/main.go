package main

import (
	"log"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/http"
	httpserver "github.com/iakigarci/go-ddd-microservice-template/pkg/http"
)

func main() {
	cfg, err := config.LoadConfig[config.Config]()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := http.NewRouter()
	server := httpserver.New(cfg, router)

	if err := <-server.Notify(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
