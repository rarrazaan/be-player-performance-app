package dependency

import (
	"context"
	"log"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	logger "github.com/rarrazaan/be-player-performance-app/internal/dependency/log"
	"github.com/rarrazaan/be-player-performance-app/internal/dependency/rdb"
	"github.com/rarrazaan/be-player-performance-app/internal/http"
	"github.com/rarrazaan/be-player-performance-app/internal/repository"
	service "github.com/rarrazaan/be-player-performance-app/internal/services"
	"github.com/rarrazaan/be-player-performance-app/internal/utils"
)

type (
	Dependency struct {
		ctx    context.Context
		Config config.Config

		Rdb *rdb.RDB
	}

	DirectDependency struct {
		Logger *logger.Logger

		*repository.Repository
		*service.Service
		*http.Handler
	}
)

func NewDependency(ctx context.Context, config config.Config) *Dependency {
	return &Dependency{
		ctx:    ctx,
		Config: config,
		Rdb:    rdb.NewRDB(ctx, config),
	}
}

func NewDirectDependency(d *Dependency) *DirectDependency {
	logger := logger.NewLogger(d.Config.LogLevel)
	jwt := utils.NewJWT(d.Config)
	repository := repository.NewRepository(d.Rdb, logger)
	service := service.NewService(repository, jwt)
	handler := http.NewHandler(service, d.Config, validator.New())

	return &DirectDependency{
		Logger:     logger,
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

func (d *Dependency) Close(ctx context.Context) <-chan struct{} {
	doneChan := make(chan struct{})

	go func() {
		shutdownOps := map[string]func(context.Context) error{
			"rdb": func(ctx context.Context) error {
				return d.Rdb.Close(ctx)
			},
		}

		wg := sync.WaitGroup{}

		for name, op := range shutdownOps {
			wg.Add(1)
			go func(name string, op func(context.Context) error) {
				defer wg.Done()
				log.Println("cleaning up:", name)
				if err := op(ctx); err != nil {
					log.Printf("failed to clean up: %s, err: %s\n", name, err)
					return
				}
				log.Println("successfully cleaned up:", name)
			}(name, op)
		}

		wg.Wait()
		close(doneChan)
	}()

	return doneChan
}
