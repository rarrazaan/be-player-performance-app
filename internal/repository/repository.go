package repository

import (
	"github.com/rarrazaan/be-player-performance-app/internal/dependency/log"
	"github.com/rarrazaan/be-player-performance-app/internal/dependency/rdb"
)

type Repository struct {
	MonoRepository IMonoRepository
}

func NewRepository(
	rdb *rdb.RDB,
	logger *log.Logger,
) *Repository {
	return &Repository{
		MonoRepository: NewMonoRepository(rdb.PostgresDB, logger),
	}
}
