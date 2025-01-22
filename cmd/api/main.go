package main

import (
	"log"

	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/http"
	httpserver "github.com/iakigarci/go-ddd-microservice-template/pkg/http"
	"github.com/iakigarci/go-ddd-microservice-template/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig[config.Config]()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.New(cfg)

	httpServer := startServers(cfg)
	if err := <-httpServer.Notify(); err != nil {
		logger.Error("Failed to start server: %v", zap.Error(err))
	}

	shutdown(httpServer, logger)
}

func startServers(cfg *config.Config) *httpserver.Server {
	router := http.NewRouter(cfg)
	server := httpserver.New(cfg, router)
	return server
}

func shutdown(server *httpserver.Server, log *zap.Logger) {
	if shutdownErr := server.Shutdown(); shutdownErr != nil {
		log.Error("app - Run - httpServer.Shutdown: %w", zap.Error(shutdownErr))
	}
}
