package di

import (
	"github.com/iakigarci/go-ddd-microservice-template/config"
	"go.uber.org/zap"
)

// Container holds all the dependencies of the application
type Container struct {
	Config *config.Config
	Logger *zap.Logger
}

func NewContainer(cfg *config.Config,
	logger *zap.Logger,
) *Container {
	return &Container{
		Config: cfg,
		Logger: logger,
	}
}

func (c *Container) Shutdown() error {

	return nil
}
